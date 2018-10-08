package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"github.com/esakat/gin-sample/db"
	"github.com/jinzhu/gorm"
	"github.com/esakat/gin-sample/config"
)


type App struct {
	db *gorm.DB
	config config.Config
}

func (app *App) Init() {
	app.db = db.InitDb()
}

func (app *App) AppendEnv(c *gin.Context) {
	log.Print("AppendEnv is executed each request.")
	c.Set("db", app.db)
}