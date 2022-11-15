package jwtToken

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

// 通过解析token创建jwt数据
func NewJwtTokenByParse(token string, key []byte) (*jwtToken, error) {
	t := &jwtToken{AccessToken: token, Key: key}
	if err := t.parse(); err != nil {
		return nil, err
	}
	return t, nil
}

// 解析
func (t *jwtToken) parse() error {
	if t.AccessToken == "" {
		return errors.New("Token无效为空")
	} else if t.Key == nil {
		return errors.New("Key无效为空")
	}

	claim, err := jwt.Parse(t.AccessToken, func(tt *jwt.Token) (interface{}, error) {
		return t.Key, nil
	})
	if err != nil {
		return err
	}

	t.Data = map[string]interface{}(claim.Claims.(jwt.MapClaims))
	return nil
}
