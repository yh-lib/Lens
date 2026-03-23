package cluster

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	logs.Info(nil, "开始运行集群更新逻辑")
	addOrUpdate(c, "update")
}
