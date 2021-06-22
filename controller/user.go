package controller

import (
	"consolidated/entity"
	"consolidated/enum"
	"consolidated/helper"
	"consolidated/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	//## Logger request
	helper.LoggerRequest(c, nil)

	user := []entity.User{}
	err := service.FindAll(&user)
	if err == nil {
		helper.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		helper.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func FindID(c *gin.Context) {
	id := c.Params.ByName("id")
	//## Logger request
	helper.LoggerRequest(c, id)

	user := entity.User{}

	err := service.FindID(&user, id)
	if err == nil {
		helper.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		helper.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func AddNewCustomer(c *gin.Context) {
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		helper.JsonResult(c, http.StatusBadRequest, enum.Success, helper.GetErrShouldBind(err))
	}

	//## Logger request
	helper.LoggerRequest(c, user)

	err = service.AddNewCustomer(&user)
	if err == nil {
		helper.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		helper.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func PutOneCustomer(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	c.BindJSON(&user)

	//## Logger request
	helper.LoggerRequest(c, user)

	err := service.PutOneCustomer(&user, id)
	if err == nil {
		helper.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		helper.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}

func DeleteCustomer(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")

	//## Logger request
	helper.LoggerRequest(c, id)

	err := service.DeleteCustomer(&user, id)
	if err == nil {
		helper.JsonResult(c, http.StatusOK, enum.Success, user)
	} else {
		helper.JsonResult(c, http.StatusInternalServerError, err.Error(), nil)
	}
}
