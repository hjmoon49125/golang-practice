package controller

import (
	"fmt"
	"net/http"
	"server_temp/source/src/model"
	"server_temp/source/src/service"

	"github.com/gin-gonic/gin"
)

type helloController struct {
	helloService model.HelloService
}

func HelloController(e *gin.RouterGroup) {
	controller := &helloController{
		helloService: service.HelloService(),
	}

	e.GET("/hello", controller.SayHello)
}

// @Summary Say Hello
// @Tags Test
// @version 1.0
// @produce application/json
// @Security jwt
// @Success 200 {object} model.Response{Data=string}
// @Router /api/hello [get]
func (h *helloController) SayHello(c *gin.Context) {
	message := h.helloService.SayHello()
	account, _ := c.Get("account")
	greeting := fmt.Sprintf(`%s %s`, message, account)
	c.JSON(http.StatusOK, model.Response{Data: greeting})
}
