package handlers

import (
	"github.com/raphaelc484/go-estudo.git/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Initialize_handler() {
	logger = config.GetLogger("handler")
	db = config.GeSQLite()
}
