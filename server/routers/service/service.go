package service

import (
	"server/controllers/service"

	"github.com/gin-gonic/gin"
)

func create(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/create", service.Create)
}
func delete(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/delete", service.Delete)
}
func deleteList(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/deleteList", service.DeleteList)
}
func update(serviceGroup *gin.RouterGroup) {
	serviceGroup.POST("/update", service.Update)
}
func get(serviceGroup *gin.RouterGroup) {
	serviceGroup.GET("/get", service.Get)
}
func list(serviceGroup *gin.RouterGroup) {
	serviceGroup.GET("/list", service.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	serviceGroup := g.Group("service")
	create(serviceGroup)
	delete(serviceGroup)
	deleteList(serviceGroup)
	update(serviceGroup)
	get(serviceGroup)
	list(serviceGroup)
}
