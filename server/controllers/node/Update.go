package node

import (
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	logs.Info(nil, "更新逻辑")
	// 定义变量
	var (
		kubeUtilser kubeutils.KubeUtilser
		returndata  config.ReturnData
		info        controllers.Info
	)
	kubeconfig := controllers.NewInfo(c, &info)
	// 通过构造函数 构建一个实例
	instance := kubeutils.NewNode(kubeconfig, nil)
	kubeUtilser = instance
	item, _ := kubeUtilser.Get("", info.Name)
	returndata.Data = map[string]any{}
	returndata.Data["item"] = item
	c.JSON(200, returndata)
}
