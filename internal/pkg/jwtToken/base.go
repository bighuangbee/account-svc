package jwtToken

import (
	"time"
)

type jwtToken struct {
	duration     time.Duration          // 有效时长 默认 1 分钟
	Data         map[string]interface{} // 数据
	AccessToken  string                 // 请求token
	RefreshToken string                 // 刷新token
	Key          []byte                 // 盐
}
