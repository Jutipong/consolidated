package handler

import (
	"consolidated/feature/auth/service"
	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

//## No Get Logger File
func Login(c *gin.Context) {
	var userRequest util.UserRequest
	code, err := service.Login(c, &userRequest)
	if err != nil {
		util.JsonResponse(c, code, "", err)
	} else {
		token := util.JwtGenerate(userRequest)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
