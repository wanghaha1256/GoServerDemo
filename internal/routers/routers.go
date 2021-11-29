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

	// GET /mds/123.md
	mdGroup.GET("/:title", mdController.GetOne)

	// PUT /mds/123.md
	mdGroup.PUT("/:title", mdController.Update)

	// POST /mds/
	mdGroup.POST("/", mdController.Create)

	// DELETE /mds/123.md
	mdGroup.DELETE(":title", mdController.Delete)

	return r

}
