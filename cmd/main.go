package main

import (
	"GoServerDemo/configs"
	"GoServerDemo/internal/util/controllers"
	"GoServerDemo/pkg"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	mdController := controllers.NewMdController()

	// server.GET("/md_lists", func(c *gin.Context) {
	// 	c.JSON(200, pkg.GetFileListWithSuffix(configs.ROOT, ".md"))
	// })

	mdGroup := server.Group("/mds")
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

	log.Fatalln(server.Run(":1234"))
}
