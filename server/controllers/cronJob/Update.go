package cronJob

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	controllers.KubectlFunc(c, "cronJob", "update")
}
