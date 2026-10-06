package config

import (
	"errors"
)

func GetPortFromConfig(config map[string]string) (string, error) {
	val, exists := config["PORT"]
	if !exists {
		return "", errors.New("ключ PORT отсутствует в конфигурации")
	}
	
	return val, nil
}