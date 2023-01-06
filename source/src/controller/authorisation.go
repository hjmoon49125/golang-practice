package controller

import (
	"net/http"
	"server_temp/source/src/middleware"
	"server_temp/source/src/model"
	"server_temp/source/src/service"

	"github.com/gin-gonic/gin"
)

type authController struct {
	accountService model.AccountService
}

func AuthController(e *gin.RouterGroup) {
	controller := &authController{
		accountService: service.AccountService(),
	}

	e.POST("/login", controller.Login)
}

// @Summary Login
// @Tags Auth
// @version 1.0
// @body application/json
// @produce application/json
// @Param request body model.AccountReq true "login params"
// @Success 200 {object} model.Response{Data=string}
// @Router /auth/login [post]
func (h *authController) Login(c *gin.Context) {
	var input model.AccountReq

	if err := c.ShouldBindJSON(&input); err != nil {
		c.String(http.StatusBadRequest, "error data")
		return
	}

	account, err := h.accountService.Login(input.Account, input.Password)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	jwtToken, _ := middleware.GenToken(account.Account)

	c.JSON(http.StatusOK, model.Response{Data: jwtToken})
}
