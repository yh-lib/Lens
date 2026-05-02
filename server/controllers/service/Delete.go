package service

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "service", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "service", "deleteList")
}
