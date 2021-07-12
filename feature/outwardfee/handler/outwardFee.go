package handler

import (
	"consolidated/feature/outwardfee/model"

	"consolidated/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutwardHandler struct{}

func (u *OutwardHandler) Branch(c *gin.Context) {
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
