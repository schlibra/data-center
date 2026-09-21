package utils

import (
	"data-center/internal/models"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() (config models.Config, err error) {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return models.Config{}, err
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	bytes, err := io.ReadAll(file)
	if err != nil {
		return models.Config{}, err
	}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return models.Config{}, err
	}
	return config, nil
}
