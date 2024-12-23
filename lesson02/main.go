package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	userGroup := r.Group("/user")
	{
		userGroup.GET("/add", func(c *gin.Context) {
			c.JSON(200, "hello world")
		})
		userGroup.GET("/delete", func(c *gin.Context) {
			c.JSON(200, "delete")
		})
		xx := userGroup.Group("/book")
		{
			xx.GET("/add", func(c *gin.Context) {
				c.JSON(200, "book-add")
			})
		}
	}
	r.Run()
}
