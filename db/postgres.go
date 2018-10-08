package db

import (
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/BurntSushi/toml"
	"github.com/esakat/gin-sample/config"
)

func InitDb() *gorm.DB {

	var config config.Config
	toml.DecodeFile("config.toml", &config)

	db, err := gorm.Open(config.DB.Driver,
		"host=" + config.DB.Host + " user=" + config.DB.User +
		" dbname=" + config.DB.Dbname + " password=" + config.DB.Password +
		" sslmode=" + config.DB.SslMode)
	if err != nil {
		panic(err)
	}

	return db
}