package handler

import (
	"consolidated/feature/inwardfee/model"

	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InwardHandler struct{}

// Iwrm godoc
// @Security ApiKeyAuth
// @Tags InwardFee
// @Accept  json
// @Produce  json
// @Param model body model.Request true "Request"
// @Success 200 {object} util.Response{responseDetail=model.Response}
// @Router /fee/inward/v1/iwrm [post]
func (u *InwardHandler) InwardFee(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResponse(c, http.StatusBadRequest, "", err.Error())
	} else {
		if code, errs := req.Validate(); errs != nil {
			util.JsonResponse(c, code, "", errs)
		} else {
			util.JsonResponse(c, code, "", model.Response{})
		}
	}
}
