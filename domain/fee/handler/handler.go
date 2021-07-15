package handler

import (
	"consolidated/base"
	"consolidated/domain/fee/model"
	"consolidated/domain/fee/service"
	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service service.IService
}

func NewHandler(service service.IService) handler {
	return handler{service: service}
}

// Fee godoc
// @Security ApiKeyAuth
// @Tags Fee
// @Accept  json
// @Produce  json
// @Param model body model.Request true "Request"
// @Success 200 {object} util.Response{responseDetail=model.Response}
// @Router /fee/v1/promotioncode [post]
func (h handler) PromotionCode(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResult(c, http.StatusBadRequest, "", err.Error())
	} else {
		result, code, err := h.service.PromotionCode(&req)
		if code == base.Successfully {
			util.JsonResult(c, code, "", result)
		} else {
			util.JsonResult(c, code, "", err)
		}
	}
}
