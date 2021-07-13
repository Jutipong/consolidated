package handler

import (
	"consolidated/feature/fee/model"

	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FeeHandler struct{}

// Fee godoc
// @Security ApiKeyAuth
// @Tags Fee
// @Accept  json
// @Produce  json
// @Param model body model.Request true "Request"
// @Success 200 {object} util.Response
// @Router /fee/v1/promotioncode [post]
func (u *FeeHandler) Fee(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResponse(c, http.StatusBadRequest, "", err.Error())
	} else {
		if code, errs := req.Validate(); errs != nil {
			util.JsonResponse(c, code, "", errs)
		} else {
			util.JsonResponse(c, code, "", req)
		}
	}
}
