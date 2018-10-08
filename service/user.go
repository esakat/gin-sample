package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/esakat/gin-sample/model"
	"github.com/jinzhu/gorm"
)

func GetUsers(c *gin.Context) {
	v, ok := c.Get("db")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	db, ok := v.(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	var users []model.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "users": users})
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	v, ok := c.Get("db")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	db, ok := v.(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	var user model.User
	if result := db.First(&user, id).RecordNotFound(); result {
		c.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user": user})
}

func PostUser(c *gin.Context) {
	var param = model.User{}
	c.Bind(&param)
	user := model.NewUser(param.Name, param.Pass)
	v, ok := c.Get("db")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	db, ok := v.(*gorm.DB)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	if dbc := db.Create(&user); dbc.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed Create"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "user": user})
}
