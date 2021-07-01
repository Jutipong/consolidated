package controller

import (
	"consolidated/features/auth/service"
	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//## No Get Logger File
func Login(c *gin.Context) {
	auth, statusCode, message, err := service.BasicAuth(c)
	if err != nil {
		utils.JsonResponse(c, statusCode, message, err)
	} else {
		token := utils.JwtGenerate(auth)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
