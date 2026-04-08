package deployment

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "deployment", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "deployment", "deleteList")
}
