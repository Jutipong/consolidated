package handler

import (
	"consolidated/enum"
	"consolidated/features/user/repository"
	"consolidated/features/user/service"
	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func (u *UserHandler) FindAll(c *gin.Context) {
	user := []repository.User{}
	err := service.FindAll(&user)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserHandler) FindID(c *gin.Context) {
	user := repository.User{}
	id := c.Params.ByName("id")

	err := service.FindID(&user, id)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserHandler) AddNewCustomer(c *gin.Context) {
	var user repository.User
	err := c.ShouldBind(&user)
	if err != nil {
		utils.JsonResult(c, http.StatusBadRequest, enum.Success, err.Error())
	}

	err = service.AddNewCustomer(&user)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserHandler) PutOneCustomer(c *gin.Context) {
	var user repository.User
	id := c.Params.ByName("id")
	c.BindJSON(&user)

	err := service.PutOneCustomer(&user, id)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserHandler) DeleteCustomer(c *gin.Context) {
	var user repository.User
	id := c.Params.ByName("id")
	err := service.DeleteCustomer(&user, id)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}
