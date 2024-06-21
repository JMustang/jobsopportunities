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
	return errors.New("Fake error!")
}

func GetLogger(p string) *Logger {
	// init logger
	logger = NewLogger(p)

	return logger
	// init db
}
