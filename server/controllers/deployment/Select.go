package deployment

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "deployment", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "deployment", "list")
}
