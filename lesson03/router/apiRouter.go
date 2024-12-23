package router

import "github.com/gin-gonic/gin"

func ApiRouterInit(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/user")
		api.GET("/news")
	}
}
