package controller

import (
	"consolidated/entity"
	"consolidated/helper"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	// custom := new(entity.Customer)
	custom := []entity.User{}
	// err := service.FindAll(&custom)
	// if err != nil {
	// 	helpers.RespondJSON(c, 404, err.Error(), custom)
	// } else {
	// 	helpers.RespondJSON(c, 200, "success", custom)
	// }

	helper.RespondJSON(c, 200, "success", custom)

	// var credentials model.Credentials
	// err := ctx.ShouldBind(&credentials)
	// if err != nil {
	// 	return ""
	// }
	// isAuthenticated := service.Login(credentials.Username, credentials.Password)
	// if isAuthenticated {
	// 	return service.GenerateToken(credentials.Username, true)
	// }
	// return ""
}
