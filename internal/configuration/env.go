package configuration

import (
	"errors"
	"os"
)

func (obj *Object) SetSecretsFromEnvironment() error {
	redisPassword := os.Getenv("TAXONOMY_REDIS_PASSWORD")
	if redisPassword == "" {
		return errors.New("TAXONOMY_REDIS_PASSWORD is required")
	}
	obj.Redis.Password = redisPassword

	return nil
}
