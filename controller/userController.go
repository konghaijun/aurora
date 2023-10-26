package controller

import (
	"auroralab/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	userService *service.UserService
}

func (ctrl *UserController) ApplyHandler(c *gin.Context) {
	newResp, err := ctrl.userService.Apply(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, newResp)
	} else {
		c.JSON(http.StatusOK, newResp)
	}

}
func (ctrl *UserController) SelectHandler(c *gin.Context) {
	newResp, err := ctrl.userService.Select(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, newResp)
	} else {
		c.JSON(http.StatusOK, newResp)
	}
}

func (ctrl *UserController) AnswerHandler(c *gin.Context) {
	newResp, err := ctrl.userService.Answer(c)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, newResp)
	} else {
		c.JSON(http.StatusOK, newResp)
	}
}
