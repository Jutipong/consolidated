package user

import (
	"consolidated/features/user/controller"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

type UserAPI struct{}

func (u *UserAPI) Setup(r *gin.Engine) {
	g := r.Group("/api/user")
	{
		user := controller.UserController{}
		g.GET("", utils.JwtVerify, user.FindAll)
		g.GET("/:id", utils.JwtVerify, user.FindID)
		g.POST("",  user.AddNewCustomer)
		g.PUT("", utils.JwtVerify, user.PutOneCustomer)
		g.DELETE("", utils.JwtVerify, user.DeleteCustomer)

	}
}
