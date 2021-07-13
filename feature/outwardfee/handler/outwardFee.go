package handler

import (
	"consolidated/feature/outwardfee/model"

	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutwardHandler struct{}

// Branch godoc
// @Security ApiKeyAuth
// @Tags Branch
// @Accept  json
// @Produce  json
// @Param model body model.Request true "Request"
// @Success 200 {object} util.Response{responseDetail=model.Response}
// @Router /fee/outward/v1/branch [post]
func (u *OutwardHandler) Branch(c *gin.Context) {
	var req model.Request
	if err := c.ShouldBind(&req); err != nil {
		util.JsonResponse(c, http.StatusBadRequest, "", err.Error())
	} else {
		if code, errs := req.Validate(); errs != nil {
			util.JsonResponse(c, code, "", errs)
		} else {
			util.JsonResponse(c, code, "", util.JsonResponse)
		}
	}
}
