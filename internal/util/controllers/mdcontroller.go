package controllers

import "github.com/gin-gonic/gin"

type mdfile struct {
	Title   string `title`
	Content string `content`
}

type MdController interface {
	GetOne(c *gin.Context)

	GetList(c *gin.Context)

	Update(c *gin.Context)

	Create(c *gin.Context)

	Delete(c *gin.Context)
}

type controller struct {
}

func (c *controller) GetOne(context *gin.Context) {

}

func (c *controller) GetList(context *gin.Context) {
	pkg.GetFileListWithSuffix(configs.ROOT, ".md")
	context.JSON("200")
}

func (c *controller) Update(context *gin.Context) {

}

func (c *controller) Create(context *gin.Context) {

}

func (c *controller) Delete(context *gin.Context) {

}

func NewMdController() MdController {
	return &controller{}
}
