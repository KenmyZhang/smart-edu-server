package api

import (
	"github.com/KenmyZhang/smart-edu-server/biz"
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateQuestion(c *gin.Context) {
	question := &model.Question{}
	if c.BindJSON(question) != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.CreateQuestion", "input params", "invalid json body"))
		return
	}

	if err := biz.CreateQuestion(question); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
		})
	}
}

func GetQuestions(c *gin.Context) {
	size := c.Query("size")
	page := c.Query("page")

	sizeInt, _ := strconv.Atoi(size)
	pageInt, _ := strconv.Atoi(page)

	if sizeInt > 50 || sizeInt <= 0 {
		sizeInt = 50
	}

	if pageInt < 0 {
		pageInt = 0
	}
	if list, err := biz.GetQuestions(sizeInt*pageInt, sizeInt); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"list": list,
		})
	}
}
