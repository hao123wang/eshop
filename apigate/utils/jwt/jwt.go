package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	tokenSecret    = []byte("这是jwt的token密钥")
	expireDuration = 24 * time.Hour // token 有效期24小时
)

// CustomCliams 自定义 token 声明
type CustomCliams struct {
	UserID   uint32 `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenToken 根据自定义 cliams 生成 token 字符串
func GenToken(userID uint32, username string) (string, error) {
	claims := &CustomCliams{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)), // 过期时间
			Issuer:    "eshop",                                            // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥签名获得完整的 token 字符串
	return token.SignedString(tokenSecret)
}

// ParseToken 从 token 字符串中解析出自定义的 claims
func ParseToken(tokenString string) (*CustomCliams, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomCliams{}, func(t *jwt.Token) (any, error) {
		return tokenSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言，并验证是否有效
	if claims, ok := token.Claims.(*CustomCliams); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
