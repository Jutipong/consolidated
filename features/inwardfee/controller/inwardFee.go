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
	} else {
		//## Validate
		ruleId, errs := req.Validate()
		if errs != nil {
			utils.JsonResult(c, ruleId, "", errs)
		} else {
			utils.JsonResult(c, float32(0000), enum.Success, req)
		}
	}
}
