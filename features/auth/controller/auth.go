package controller

import (
	"consolidated/features/auth/service"
	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//## No Get Logger File
func Login(c *gin.Context) {
	auth, err := service.BasicAuth(c)
	if err != nil {
		utils.JsonResult(c, http.StatusUnauthorized, "", err)
		// utils.JsonResult(c, http.StatusUnauthorized, "username or password is not authenticated", nil)
	} else {
		//invalid in db
		token := utils.JwtGenerate(auth)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
