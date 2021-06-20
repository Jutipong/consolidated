package controller

import (
	"consolidated/entity"
	"consolidated/enum"
	"consolidated/helper"
	"consolidated/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindID(c *gin.Context) {
	id := c.Params.ByName("id")
	user := entity.User{}

	err := service.FindID(&user, id)
	if err == nil {
		helper.RespondJSON(c, http.StatusOK, enum.Success, user)
	} else {
		helper.RespondJSON(c, http.StatusInternalServerError, err.Error(), user)
	}
}

func FindAll(c *gin.Context) {
	user := []entity.User{}

	err := service.FindAll(&user)
	if err == nil {
		helper.RespondJSON(c, http.StatusOK, enum.Success, user)
	} else {
		helper.RespondJSON(c, http.StatusInternalServerError, err.Error(), user)
	}
}

func AddNewCustomer(c *gin.Context) {
	var user entity.User

	err := c.ShouldBind(&user)
	if err != nil {
		helper.RespondJSON(c, http.StatusBadRequest, enum.Success, helper.GetErrShouldBind(err))
	}

	c.BindJSON(&user)
	err = service.AddNewCustomer(&user)
	if err == nil {
		helper.RespondJSON(c, http.StatusOK, enum.Success, user)
	} else {
		helper.RespondJSON(c, http.StatusInternalServerError, err.Error(), user)
	}
}

func PutOneCustomer(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	c.BindJSON(&user)

	err := service.PutOneCustomer(&user, id)
	if err == nil {
		helper.RespondJSON(c, http.StatusOK, enum.Success, user)
	} else {
		helper.RespondJSON(c, http.StatusInternalServerError, err.Error(), user)
	}
}

func DeleteCustomer(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")

	err := service.DeleteCustomer(&user, id)
	if err == nil {
		helper.RespondJSON(c, http.StatusOK, enum.Success, user)
	} else {
		helper.RespondJSON(c, http.StatusInternalServerError, err.Error(), user)
	}
}
