package controllers

import (
	"GoServerDemo/internal/util/models"

	"github.com/gin-gonic/gin"
)

type mdfile struct {
	models.Model

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	CreateOn   string `json:""`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`

	// TODO: 创建时间和修改时间
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
