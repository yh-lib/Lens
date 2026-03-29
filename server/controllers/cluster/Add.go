package cluster

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	logs.Info(nil, "集群添加逻辑")
	addOrUpdate(c, "add")
}
