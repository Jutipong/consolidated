package handler

import (
	"consolidated/domain/fee/model"

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
// @Success 200 {object} util.Response{responseDetail=model.Response}
// @Router /fee/v1/promotioncode [post]
func (u *FeeHandler) Fee(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResult(c, http.StatusBadRequest, "", err.Error())
	} else {
		if code, err := req.Validate(); err != nil {
			util.JsonResult(c, code, "", err)
		} else {
			util.JsonResult(c, code, "", model.Response{})
		}
	}
}
