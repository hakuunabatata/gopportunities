package config

import (
	gorm "gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	return nil
}
