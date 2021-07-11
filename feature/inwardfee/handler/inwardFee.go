package handler

import (
	"consolidated/base"
	"consolidated/feature/inwardfee/model"

	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InwardHandler struct{}

func (u *InwardHandler) InwardFee(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResponse(c, http.StatusBadRequest, base.Successfully, err.Error())
	} else {
		if code, errs := req.Validate(); errs != nil {
			util.JsonResponse(c, code, "", errs)
		} else {
			util.JsonResponse(c, code, base.Successfully, req)
		}
	}
}
