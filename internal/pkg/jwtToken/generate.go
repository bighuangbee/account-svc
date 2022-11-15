package jwtToken

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 通过生成token创建jwt数据
// d 有效时长   key 盐  data 数据
func NewJwtTokenByGenerate(d time.Duration, key []byte, data map[string]interface{}) (*jwtToken, error) {
	t := &jwtToken{
		duration: d,
		Key:      key,
		Data:     data,
	}
	if err := t.generate(); err != nil {
		return nil, err
	}
	return t, nil
}

// 生成token
func (t *jwtToken) generate() error {
	if t.AccessToken != "" {
		return nil
	}

	t.initDefault()
	if t.Key == nil {
		return errors.New("key无效为空")
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(t.Data))
	if token, err := at.SignedString(t.Key); err != nil {
		return err
	} else {
		t.AccessToken = token
	}

	return nil
}

// 初始化参数
func (t *jwtToken) initDefault() {
	if t.Data == nil {
		t.Data = map[string]interface{}{}
	}
	if t.duration == 0 {
		t.duration = 2 * time.Hour
	}
	// 设置过期时间
	t.Data["exp"] = time.Now().Add(t.duration).Unix()
}
