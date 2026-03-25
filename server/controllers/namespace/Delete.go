package namespace

import (
	"context"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(c *gin.Context) {
	logs.Info(nil, "删除逻辑")
	returnData := config.ReturnData{}
	clientSet, basicInfo, err := controllers.BasicInit(c)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
	}
	var nameSpace corev1.Namespace
	nameSpace.Name = basicInfo.NameSpace
	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), nameSpace.Name, metav1.DeleteOptions{})
	if err != nil {
		returnData.Status = 400
		returnData.Message = "删除 namespace: " + nameSpace.Name + " 失败:" + err.Error()
		logs.Error(map[string]any{"Error": err}, "删除 namespace: "+nameSpace.Name+" 失败:")
	} else {
		returnData.Status = 200
		returnData.Message = "删除 namespace: " + nameSpace.Name + " 成功:"
	}
	c.JSON(200, returnData)
}
