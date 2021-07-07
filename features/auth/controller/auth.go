package controller

import (
	"consolidated/features/auth/model"
	"consolidated/features/auth/service"
	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//## No Get Logger File
func Login(c *gin.Context) {
	var userRequest model.UserRequest
	code, err := service.Login(c,&userRequest)
	if err != nil {
		utils.JsonResult(c, code, "", err)
	} else {
		token := utils.JwtGenerate(userRequest)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
