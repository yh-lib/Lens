package cronJob

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "cronJob", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "cronJob", "list")
}
