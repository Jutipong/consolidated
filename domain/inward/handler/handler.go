package handler

import (
	"consolidated/base"
	"consolidated/domain/inward/model"
	"consolidated/domain/inward/service"

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

// Iwrm godoc
// @Security ApiKeyAuth
// @Tags InwardFee
// @Accept  json
// @Produce  json
// @Param model body model.Request true "Request"
// @Success 200 {object} util.Response{responseDetail=model.Response}
// @Router /fee/inward/v1/iwrm [post]
func (h handler) Iwrm(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResult(c, http.StatusBadRequest, "", err.Error())
	} else {
		result, code, err := h.service.Iwrm(&req)
		if code == base.Successfully {
			util.JsonResult(c, code, "", result)
		} else {
			util.JsonResult(c, code, "", err)
		}
	}
}
