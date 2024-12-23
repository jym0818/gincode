package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jym/gincode/lesson03/router"
)

func main() {
	r := gin.Default()
	router.AdminRouterInit(r)
	router.ApiRouterInit(r)
	r.Run()
}
