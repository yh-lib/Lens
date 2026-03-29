// 控制器层  实现路由的处理逻辑
package controllers

import (
	"errors"
	"server/config"
	"server/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type BasicInfo struct {
	ClusterId  string   `json:"clusterId" form:"clusterId"`
	NameSpace  string   `json:"nameSpace" form:"nameSpace"`
	Name       string   `json:"name" form:"name"`
	Item       any      `json:"item"`
	DeleteList []string `json:"deleteList"`
}

type Info struct {
	BasicInfo
	ReturnData    config.ReturnData
	LabelSelector string `json:"labelSelector" form:"labelSelector"`
	FieldSelector string `json:"fieldSelector" form:"fieldSelector"`
	// 判断是否是强制删除
	Force bool `json:"force" form:"force"`
}

func NewInfo(c *gin.Context, info *Info) (kubeconfig string) {
	var (
		err        error
		returnData config.ReturnData
		// 赋值变量部分
		requestMethod = c.Request.Method
	)
	// 前后端数据绑定
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&info)
	case "POST":
		err = c.ShouldBindJSON(&info)
	default:
		err = errors.New("不支持的请求类型")
	}
	logs.Debug(map[string]any{"Info": info}, "数据绑定结果")
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(200, returnData)
		logs.Error(nil, msg)
		return
	}
	// 获取kubeconfig
	return config.ClusterKubeconfig[info.ClusterId]
}

func BasicInit(c *gin.Context, item any) (clientSet *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	// 绑定前端传递的数据到后端变量中
	basicInfo = BasicInfo{}
	basicInfo.Item = item
	requestMethod := c.Request.Method
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&basicInfo)
	case "POST":
		err = c.ShouldBindJSON(&basicInfo)
	default:
		err = errors.New("不支持的请求类型")
	}
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.NameSpace == "" {
		basicInfo.NameSpace = "default"
	}
	// 实例化clientSet
	kubeConfig := config.ClusterKubeconfig[basicInfo.ClusterId]
	kubeConfigByte := []byte(kubeConfig)
	// kubeConfig
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfigByte)
	if err != nil {
		msg := "实例化 kubeconfig 失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// clientSet
	clientSet, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "实例化clientSet失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// 校验 kubeconfig
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		msg := "kubeconfig 校验失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	logs.Debug(nil, "集群版本："+serverVersion.String())
	return clientSet, basicInfo, nil
}

func newReturnData(c *gin.Context, err error, returndata config.ReturnData, msg string) {
	logs.Error(map[string]any{"ERROR": err.Error()}, msg)
	returndata.Status = 200
	returndata.Message = msg + ": " + err.Error()
	c.JSON(200, returndata)
}

func KubectlFunc(c *gin.Context, resourceType string, opMethod string) {
	logs.Debug(nil, resourceType+"列表逻辑")
	// 定义变量
	var (
		kubeUtilser kubeutils.KubeUtilser
		returndata  config.ReturnData
		info        Info
		item        corev1.Node
	)
	// 初始化返回数据
	returndata.Data = map[string]any{}
	info.Item = &item
	kubeconfig := NewInfo(c, &info)
	// 匹配资源类型
	switch resourceType {
	case "node":
		kubeUtilser = kubeutils.NewNode(kubeconfig, &item)
	default:
		logs.Error(nil, "不支持该资源类型")
		return
	}
	// 匹配操作方法
	switch opMethod {
	case "list":
		items, err := kubeUtilser.List("", info.LabelSelector, info.FieldSelector)
		if err != nil {
			newReturnData(c, err, returndata, "获取列表失败")
			return
		}
		returndata.Data["items"] = items
	case "get":
		item, err := kubeUtilser.Get("", info.Name)
		if err != nil {
			newReturnData(c, err, returndata, "获取详情失败")
			return
		}
		returndata.Data["items"] = item
	case "update":
		err := kubeUtilser.Update(info.NameSpace)
		if err != nil {
			newReturnData(c, err, returndata, "更新失败2")
			return
		}
	default:
		logs.Error(nil, "不支持该操作方法")
		return
	}
	returndata.Status = 200
	returndata.Message = "操作成功"
	c.JSON(200, returndata)
}
