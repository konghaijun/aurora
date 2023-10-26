package router

import (
	"auroralab/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	userCtrl := &controller.UserController{}
	// 添加跨域中间件
	router.Use(CORSMiddleware())
	router.POST("/aurora/user/apply", userCtrl.ApplyHandler)
	router.GET("/aurora/user/select", userCtrl.SelectHandler)
	router.POST("/aurora/user/chat", userCtrl.AnswerHandler)
	return router
}

// 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
