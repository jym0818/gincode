package router

import (
	"github.com/gin-gonic/gin"
	admin2 "github.com/jym/gincode/lesson03/controller/admin"
)

func AdminRouterInit(r *gin.Engine) {
	admin := r.Group("/admin")
	{
		admin.GET("/user", (&admin2.UserController{}).UserIndex)
		admin.GET("/news", (&admin2.NewsController{}).NewsIndex)
	}
}
