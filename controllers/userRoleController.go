package controllers

import (
	"booking-api/data/response"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserRoleController struct {
	UserRoleService services.UserRoleService
}

func NewUserRoleController(userRoleService services.UserRoleService) *UserRoleController {
	return &UserRoleController{UserRoleService: userRoleService}
}

func (u *UserRoleController) FindAll(ctx *gin.Context) {

	userRoles := u.UserRoleService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userRoles,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (u *UserRoleController) FindByID(ctx *gin.Context) {
	userRoleID := ctx.Param("userRoleID")

	userRoleIDInt, _ := strconv.Atoi(userRoleID)

	userRole := u.UserRoleService.FindByID(uint(userRoleIDInt))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   userRole,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (u *UserRoleController) Create(ctx *gin.Context) {

	var createUserRoleRequest string

	err := ctx.ShouldBindJSON(&createUserRoleRequest)

	if err != nil {
		panic(err)
	}

	result := u.UserRoleService.Create(createUserRoleRequest)

	webResponse := response.Response{
		Code:   201,
		Status: "Created",
		Data:   result,
	}

	ctx.JSON(http.StatusCreated, webResponse)

}
