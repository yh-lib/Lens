package cronJob

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "cronJob", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "cronJob", "deleteList")
}
