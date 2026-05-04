package cronJob

import (
	"server/controllers/cronJob"

	"github.com/gin-gonic/gin"
)

func create(cronJobGroup *gin.RouterGroup) {
	cronJobGroup.POST("/create", cronJob.Create)
}
func delete(cronJobGroup *gin.RouterGroup) {
	cronJobGroup.POST("/delete", cronJob.Delete)
}
func deleteList(cronJobGroup *gin.RouterGroup) {
	cronJobGroup.POST("/deleteList", cronJob.DeleteList)
}
func update(cronJobGroup *gin.RouterGroup) {
	cronJobGroup.POST("/update", cronJob.Update)
}
func get(cronJobGroup *gin.RouterGroup) {
	cronJobGroup.GET("/get", cronJob.Get)
}
func list(cronJobGroup *gin.RouterGroup) {
	cronJobGroup.GET("/list", cronJob.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	cronJobGroup := g.Group("cronJob")
	create(cronJobGroup)
	delete(cronJobGroup)
	deleteList(cronJobGroup)
	update(cronJobGroup)
	get(cronJobGroup)
	list(cronJobGroup)
}
