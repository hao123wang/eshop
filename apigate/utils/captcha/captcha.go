package captcha

import (
	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type CaptchaMana struct {
	Store base64Captcha.Store
}

func NewCaptchaMana(rdb *redis.Client) *CaptchaMana {
	return &CaptchaMana{
		Store: NewRedisStore("captcha:", rdb),
	}
}

func (m *CaptchaMana) Generate() (string, string, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, m.Store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		return "", "", err
	}
	return id, b64s, nil
}

func (m *CaptchaMana) Verify(id, captchaStr string) bool {
	return m.Store.Verify(id, captchaStr, true)
}
