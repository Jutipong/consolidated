package controller

import (
	"consolidated/helper"
	"consolidated/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var login model.Login
	err := c.ShouldBind(&login)
	if err != nil {
		helper.RespondJSON(c, http.StatusUnauthorized,
			"username or password is not authenticated",
			helper.GetErrShouldBind(err))
	} else {
		token := helper.JwtGenerate(login)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
