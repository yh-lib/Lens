package deployment

import (
	"server/controllers/deployment"

	"github.com/gin-gonic/gin"
)

func create(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/create", deployment.Create)
}
func delete(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/delete", deployment.Delete)
}
func deleteList(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/deleteList", deployment.DeleteList)
}
func update(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.POST("/update", deployment.Update)
}
func get(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.GET("/get", deployment.Get)
}
func list(deploymentGroup *gin.RouterGroup) {
	deploymentGroup.GET("/list", deployment.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	deploymentGroup := g.Group("deployment")
	create(deploymentGroup)
	delete(deploymentGroup)
	deleteList(deploymentGroup)
	update(deploymentGroup)
	get(deploymentGroup)
	list(deploymentGroup)
}
