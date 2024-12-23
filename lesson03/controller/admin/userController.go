package admin

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (con *UserController) UserIndex(c *gin.Context) {
	c.JSON(200, "admin-user")
}
