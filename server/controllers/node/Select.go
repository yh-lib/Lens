package node

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "node", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "node", "list")
}
