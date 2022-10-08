package routers

import (
	"gin-blog/controllers"
	"github.com/gin-gonic/gin"
)

/**
 * @Description
 * @Author yingtie
 * @Date 2022/9/30 11:44
 **/
func init() {
	r := gin.Default()
	r.GET("/", controllers.Hello)
	r.GET("/ping", controllers.Ping)
	r.Run(":8080")
}
