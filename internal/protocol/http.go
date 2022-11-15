package protocol

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"time"
	netHttp "net/http"
	kjson "github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/bighuangbee/account-svc/internal/conf"
	"github.com/bighuangbee/account-svc/internal/data"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(bc *conf.Bootstrap, logger log.Logger, server *PbServer, bundle *i18n.Bundle, data *data.Data) *http.Server {
	// 不需要验证token的地址
	//checkTokenWhiteList := []string{
	//	"/api.mozi.device.v1.Device/SyncWvp",
	//	"/api.mozi.device.v1.Device/DeviceRecordHook",
	//}
	c := bc.Server
	srv := http.NewServer(
		http.Address(c.HTTP.Addr),
		http.Timeout(time.Duration(c.HTTP.Timeout)*time.Second),
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			//hiGoi18n.Translator(bundle),
			validate(),
			//opLogCli.SaveOpLog(),
			//hiKratos.HTTPReturnTraceID(),
		),

		http.Logger(logger),
		// hiKratos.Encoder(),
		// hiKratos.ErrorEncoder(),

		Encoder(),
		ErrorEncoder(),
	)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	server.RegisterHTTP(srv)

	return srv
}

type MyError struct {
	Code      int32             `json:"code,omitempty"`
	DetailMsg string            `json:"detailMsg,omitempty"`
	Message   string            `json:"message,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
	RetCode   int32             `json:"retCode"`
}

func ErrorEncoder() http.ServerOption {
	return http.ErrorEncoder(func(w netHttp.ResponseWriter, r *netHttp.Request, err error) {
		// 拿到error并转换成kratos Error实体
		se := errors.FromError(err)

		// 通过Request Header的Accept中提取出对应的编码器
		codec, _ := http.CodecForRequest(r, "Accept")
		body, err := codec.Marshal(&MyError{
			DetailMsg: se.Reason,
			Message:   se.Message,
			Metadata:  se.Metadata,
			RetCode:   se.Code,
		})
		if err != nil {
			w.WriteHeader(netHttp.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		code := netHttp.StatusOK
		if se.Code == netHttp.StatusUnauthorized {
			code = netHttp.StatusUnauthorized
		}
		w.WriteHeader(code)
		w.Write(body)
		return
	})
}

const (
	defaultReason  = "SUCCESS"
	defaultMessage = "操作成功"
)

func Encoder() http.ServerOption {
	return http.ResponseEncoder(func(w netHttp.ResponseWriter, r *netHttp.Request, v interface{}) error {
		msg := GetMessage(w)
		if msg == "" {
			msg = defaultMessage
		}
		codec, _ := http.CodecForRequest(r, "Accept")
		// 枚举使用数字
		kjson.MarshalOptions.UseEnumNumbers = true
		data, err := codec.Marshal(v)


		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")

		_, _ = w.Write([]byte(fmt.Sprintf(`{
			"code": %d,
			"detailMsg": "%s",
			"message": "%s",
			"data": %s
			}`, 0, defaultReason, msg, data)))
		return nil
	})
}

var headerKey = "_response-msg_"

// GetMessage 重写http.ResponseEncoder时，获取SetMessage设置的值
func GetMessage(w netHttp.ResponseWriter) string {
	msg := w.Header().Get(headerKey)
	if msg != "" {
		w.Header().Del(headerKey)
	}
	return msg
}
