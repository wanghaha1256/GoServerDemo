package controllers

import (
	"GoServerDemo/configs"
	"GoServerDemo/internal/util/models"
	"GoServerDemo/internal/util/settings"
	"GoServerDemo/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

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
	id := com.StrTo(context.Param("id")).MustInt()

	statusCode := configs.BAD_REQUEST
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
		statusCode = configs.SUCCESS
	}

	context.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"data": data,
	})
}

func (c *controller) GetList(context *gin.Context) {

	fmt.Printf("%+v", context)

	data := make(map[string]interface{})
	keys := "title LIKE ?"
	values := "%" + context.Query("title") + "%"

	statusCode := configs.SUCCESS

	data["lists"] = models.GetMdList(pkg.GetPage(context), settings.PageSize, keys, values)
	data["total"] = models.GetMdNum(keys, values)

	context.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"data": data,
	})
}

func (c *controller) Update(context *gin.Context) {

	id := com.StrTo(context.Param("id")).MustInt()
	title := context.PostForm("title")
	desc := context.PostForm("desc")
	content := context.PostForm("content")
	modifiedBy := context.PostForm("modified_by")

	statusCode := configs.NOT_FOUND

	if models.ExistArticleByID(id) {
		data := make(map[string]interface{})
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["modified_by"] = modifiedBy

		models.UpdateArticle(id, data)
		statusCode = configs.SUCCESS
	}

	context.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"data": make(map[string]interface{}),
	})
}

func (c *controller) Create(context *gin.Context) {

	fmt.Printf("%+v", context)

	title := context.PostForm("title")
	desc := context.PostForm("desc")
	content := context.PostForm("content")
	createdBy := context.PostForm("created_by")

	statusCode := configs.BAD_REQUEST

	data := make(map[string]interface{})
	data["title"] = title
	data["desc"] = desc
	data["content"] = content
	data["created_by"] = createdBy

	models.AddArticle(data)
	statusCode = configs.SUCCESS

	context.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"data": make(map[string]interface{}),
	})

}

func (c *controller) Delete(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()

	statusCode := configs.NOT_FOUND
	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
		statusCode = configs.SUCCESS
	}

	context.JSON(http.StatusOK, gin.H{
		"code": statusCode,
		"data": make(map[string]interface{}),
	})
}

func NewMdController() MdController {
	return &controller{}
}
