package config

import (
	"fmt"

	"github.com/JoshEvan/alterra-agmc-day4/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type configDB struct {
	username string
	password string
	port     string
	host     string
	database string
}

func InitDB() {
	config := configDB{
		username: "root",
		password: "temp123",
		port:     "3306",
		host:     "127.0.0.1",
		database: "agmcday2",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.username,
		config.password,
		config.host,
		config.port,
		config.database,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	initMigrate()
}

func initMigrate() {
	DB.AutoMigrate(&models.User{})
}
