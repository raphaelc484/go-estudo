package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	return errors.New("fake error")
}

func GetLogger(p string) *Logger {
	logger = New_logger(p)
	return logger
}
