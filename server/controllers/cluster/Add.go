package cluster

import (
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	clusterID := c.Query("clusterid")
	logs.Info(nil, "开始运行集群 "+clusterID+" 添加逻辑")
	addOrUpdate(c, "create")
}
