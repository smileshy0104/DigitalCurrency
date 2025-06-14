package tools

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// ParseToken 解析JWT令牌，验证其有效性并返回用户ID和过期时间
// 参数:
//
//	tokenString: JWT令牌字符串
//	secret: 用于签名验证的密钥
//
// 返回值:
//
//	int64: 用户ID
//	error: 错误信息，如果令牌无效或解析出错
func ParseToken(tokenString string, secret string) (int64, error) {
	// 解析JWT令牌，使用指定的密钥进行签名验证
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否为预期的HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret 是一个包含密钥的 []byte，例如 []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	// 验证令牌的声明和有效性
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 提取并转换用户ID和过期时间
		val := claims["userId"].(float64)
		exp := int64(claims["exp"].(float64))
		// 检查令牌是否已过期
		if exp <= time.Now().Unix() {
			return 0, errors.New("token过期了")
		}
		return int64(val), nil
	} else {
		return 0, err
	}
}
