package controller

import (
	"consolidated/entity"
	"consolidated/helper"
	"consolidated/service"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			helper.RespondJSON(c, 404, "errors", helper.Descriptive(verr))
			return
		}
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
