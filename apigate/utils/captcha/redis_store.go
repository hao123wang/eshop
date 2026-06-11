package captcha

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
)

type RedisStore struct {
	RDB       *redis.Client
	KeyPrefix string
}

func NewRedisStore(keyPrefix string, rdb *redis.Client) *RedisStore {
	return &RedisStore{
		RDB:       rdb,
		KeyPrefix: keyPrefix,
	}
}

var _ base64Captcha.Store = (*RedisStore)(nil)

func (s *RedisStore) Set(id string, value string) error {
	key := s.KeyPrefix + id
	return s.RDB.Set(key, value, 5*time.Minute).Err()
}

func (s *RedisStore) Get(id string, clear bool) string {
	key := s.KeyPrefix + id
	v, err := s.RDB.Get(key).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.RDB.Del(key)
	}
	return v
}

func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
