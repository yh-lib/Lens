package service

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "service", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "service", "list")
}
