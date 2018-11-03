package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"smart-edu-server/biz"
)

func GetArticle(c *gin.Context) {
	articleId := c.Param("article_id")
	if data, err := biz.GetArticle(articleId); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	} else {
		c.Header("content-disposition", `attachment; filename=`+articleId)
		c.Data(200, `Content-Type: text/md; charset=utf-8`, data)
		c.Writer.Write(data)
	}
}
