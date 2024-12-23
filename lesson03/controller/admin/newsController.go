package admin

import "github.com/gin-gonic/gin"

type NewsController struct {
}

func (con *NewsController) NewsIndex(c *gin.Context) {
	c.JSON(200, "admin-news")
}
