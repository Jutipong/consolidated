package handler

import (
	"consolidated/enum"
	"consolidated/features/outwardfee/model"

	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutwardHandler struct{}

func (u *OutwardHandler) Branch(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		utils.JsonResult(c, http.StatusBadRequest, enum.Success, err.Error())
	} else {
		if code, errs := req.Validate(); errs != nil {
			utils.JsonResult(c, code, "", errs)
		} else {
			utils.JsonResult(c, code, enum.Success, req)
		}
	}
}
