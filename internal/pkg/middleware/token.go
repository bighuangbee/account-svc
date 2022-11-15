package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bighuangbee/gokit/kitKratos"
	pb "github.com/bighuangbee/account-svc/api/account/v1"
	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/bighuangbee/account-svc/internal/data"
	"github.com/bighuangbee/account-svc/internal/pkg/jwtToken"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

var (
	jwtTokenKey  = []byte("bighuangbee")
	userTokenKey = "account:token"
)

type (
	JwtToken     struct{}
	RefreshToken struct{}
	RequestType  struct{ Type string }
)


// CheckTokenMiddleWare Check Token middleware
func CheckTokenMiddleWare(bc *conf.Bootstrap, data *data.Data) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				if token := tr.RequestHeader().Get("jwtToken"); token != "" {
					jwtToken, err := jwtToken.NewJwtTokenByParse(token, jwtTokenKey)
					if err != nil {
						return nil, kitKratos.ResponseErr(ctx, pb.ErrorUnauthenticated)
					}
					// 校验最新token, 互踢
					cli := data.Redis("")
					if cli != nil {
						userId := jwtToken.Data["userId"]
						key := fmt.Sprintf("%s:%.0f", userTokenKey, userId)
						exist, err := cli.RedisExist(key)
						if err != nil {
							fmt.Println("jwt redis err,", err)
							return nil, kitKratos.ResponseErr(ctx, pb.ErrorInternalError)
						}
						if exist > 0 {
							oldToken, err := cli.RedisGet(key)
							if err != nil {
								fmt.Println("jwt redis err,", err)
								return nil, kitKratos.ResponseErr(ctx, pb.ErrorInternalError, "redis err: %s", err.Error())
							}
							if oldToken == token {
								ctx = context.WithValue(ctx, JwtToken{}, jwtToken.Data)
								return handler(ctx, req)
							}
						}
					}
					return nil, kitKratos.ResponseErr(ctx, pb.ErrorUnauthenticated)
				}
			}
			return nil, pb.ErrorUnauthenticated("CheckToken failed.")
		}
	}
}

// grpc,注册空context值
func CheckTokenMiddleWareGrpc() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			ctx = context.WithValue(ctx, RequestType{}, RequestType{Type: "grpc"})
			return handler(ctx, req)
		}
	}
}

type TokenData struct {
	Account  string   `json:"account"`
	UserID   uint32   `json:"userId"`
	UserName string   `json:"userName"`
	CorpId   string   `json:"corpId"`
	GroupId  []string `json:"groupId"`
	Token    string   `json:"token"`
	jwt.StandardClaims
}

// 返回 token, 是否是http,err
func GetCtxToken(ctx context.Context) (*TokenData, error) {
	jwtToken := ctx.Value(JwtToken{})
	if jwtToken == nil {
		return nil, errors.New("没有检测到token数据")
	}
	if data, ok := jwtToken.(map[string]interface{}); ok {
		reply := &TokenData{}
		dataByte, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(dataByte, reply)
		if err != nil {
			return nil, err
		}
		// http请求
		return reply, nil
	}
	return nil, errors.New("token解析错误")
}

// 返回 token, 是否是http,err
func GetUserInfo(ctx context.Context, reqUserId uint64, reqUserName string) (userId uint64, userName string, err error) {
	reqType := ctx.Value(RequestType{})
	if reqType != nil {
		if v, ok := reqType.(RequestType); ok && v.Type == "grpc" {
			// 来自grpc请求
			return reqUserId, reqUserName, nil
		}
	}

	// 来自http请求
	token, err := GetCtxToken(ctx)
	if err != nil {
		return 0, "", err
	} else {
		return uint64(token.UserID), token.UserName, nil
	}
}
