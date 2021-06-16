package controller

import (
	"consolidated/entity"
	"consolidated/helper"
	"consolidated/service"

	"github.com/gin-gonic/gin"
)

func FindID(c *gin.Context) {
	id := c.Params.ByName("id")
	user := entity.User{}
	err := service.FindID(&user, id)
	if err != nil {
		helper.RespondJSON(c, 404, err.Error(), user)
	} else {
		helper.RespondJSON(c, 200, "success", user)
	}
}

func FindAll(c *gin.Context) {
	user := []entity.User{}
	err := service.FindAll(&user)
	if err != nil {
		helper.RespondJSON(c, 404, err.Error(), user)
	} else {
		helper.RespondJSON(c, 200, "success", user)
	}
}

func AddNewCustomer(c *gin.Context) {
	var user entity.User
	err := c.ShouldBind(&user)
	c.BindJSON(&user)
	if err != nil {
		helper.RespondJSON(c, 404, "errors", helper.GetErrShouldBind(err))
	}
	err = service.AddNewCustomer(&user)
	if err != nil {
		helper.RespondJSON(c, 404, err.Error(), user)
	} else {
		helper.RespondJSON(c, 200, "success", user)
	}
}

func PutOneCustomer(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	err := service.FindID(&user, id)
	if err != nil {
		helper.RespondJSON(c, 404, err.Error(), user)
	}
	c.BindJSON(&user)
	err = service.PutOneCustomer(&user, id)
	if err != nil {
		helper.RespondJSON(c, 404, err.Error(), user)
	} else {
		helper.RespondJSON(c, 200, "success", user)
	}
}

func DeleteCustomer(c *gin.Context) {
	var user entity.User
	id := c.Params.ByName("id")
	err := service.DeleteCustomer(&user, id)
	if err != nil {
		helper.RespondJSON(c, 404, err.Error(), user)
	} else {
		helper.RespondJSON(c, 200, "success", user)
	}
}
