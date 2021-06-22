package controller

import (
	"consolidated/helper"
	"consolidated/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticatedController struct{}

//## No Get Logger File
func (auth *AuthenticatedController) Login(c *gin.Context) {
	var login model.Login
	err := c.ShouldBind(&login)
	if err != nil {
		helper.JsonResult(c, http.StatusUnauthorized,
			"username or password is not authenticated",
			helper.GetErrShouldBind(err))
	} else {
		token := helper.JwtGenerate(login)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
