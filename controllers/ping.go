package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @Description
 * @Author yingtie
 * @Date 2022/9/30 11:25
 **/
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome gin",
		"date":    time.Now(),
	})
}
