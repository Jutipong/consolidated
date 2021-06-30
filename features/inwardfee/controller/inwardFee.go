package controller

import (
	"consolidated/enum"
	"consolidated/features/inwardfee/model"

	// "consolidated/features/inwardfee/service"
	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var _inwardFee = service.InwardFee{}

type UserController struct{}

func (u *UserController) InwardFee(c *gin.Context) {
	var req model.Request
	err := c.ShouldBind(&req)
	if err != nil {
		utils.JsonResult(c, http.StatusBadRequest, enum.Success, err.Error())
	}

	//## Validate
	statusCode, message, errs := req.Validate()
	if errs != nil {
		utils.JsonResponse(c, statusCode, message, errs)
	}

	// req := []model.Request{}
	// err := _inwardFee.FindById(&req)
	// if err == nil {
	// 	utils.JsonResult(c, http.StatusOK, enum.Success, req)
	// } else {
	// 	utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	// }
}
