package controller

import (
	"consolidated/enum"
	"consolidated/features/user/repository"
	"consolidated/features/user/service"
	"consolidated/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) FindAll(c *gin.Context) {
	user := []repository.User{}
	err := service.FindAll(&user)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserController) FindID(c *gin.Context) {
	user := repository.User{}
	id := c.Params.ByName("id")

	err := service.FindID(&user, id)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserController) AddNewCustomer(c *gin.Context) {
	var user repository.User
	err := c.ShouldBind(&user)
	if err != nil {
		utils.JsonResult(c, http.StatusBadRequest, enum.Success, utils.GetErrShouldBind(err))
	}

	err = service.AddNewCustomer(&user)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (u *UserController) PutOneCustomer(c *gin.Context) {
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

func (u *UserController) DeleteCustomer(c *gin.Context) {
	var user repository.User
	id := c.Params.ByName("id")
	err := service.DeleteCustomer(&user, id)
	if err == nil {
		utils.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		utils.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}
