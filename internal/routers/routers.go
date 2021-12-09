package routers

import (
	"GoServerDemo/internal/util/controllers"
	"GoServerDemo/internal/util/settings"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(settings.RunMode)

	mdController := controllers.NewMdController()

	// server.GET("/md_lists", func(c *gin.Context) {
	// 	c.JSON(200, pkg.GetFileListWithSuffix(configs.ROOT, ".md"))
	// })

	mdGroup := r.Group("/mds")
	// GET /md_lists
	mdGroup.GET("/md_lists", mdController.GetList)

	// GET /mds/1
	mdGroup.GET("/:id", mdController.GetOne)

	// PUT /mds/1
	mdGroup.POST("/:id", mdController.Update)

	// POST /mds/
	mdGroup.POST("/create", mdController.Create)

	// DELETE /mds/1
	mdGroup.DELETE(":id", mdController.Delete)

	return r

}
