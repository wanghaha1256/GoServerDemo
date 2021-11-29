package util

import (
	"GoServerDemo/configs"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(context *gin.Context) int {
	result := 0
	page, err := com.StrTo(context.Query("page")).Int()
	if err != nil {
		log.Fatalln(err)
		page = 1
	}
	if page > 0 {
		result = (page - 1) * configs.BAD_REQUEST
	}
	return result
}
