package controller

import (
	"consolidated/helper"
	"consolidated/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var login model.Login
	err := ctx.ShouldBind(&login)
	if err != nil {
		helper.RespondJSON(ctx, http.StatusUnauthorized,
			"username or password is not authenticated",
			helper.GetErrShouldBind(err))
	} else {
		token := helper.JwtGenerate(login)
		helper.RespondJSON(ctx, http.StatusOK, "login ok", gin.H{"token": token})
	}
}
