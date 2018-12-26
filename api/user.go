package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/KenmyZhang/smart-edu-server/biz"
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/model"
)

func GetUser(c *gin.Context) {
	userId := c.Param("user_id")
	if user, err := biz.GetUser(userId); err != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.GetUser", "User-ID", err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"result": "ok",
			"user":   user,
		})
	}
}

func CreateUser(c *gin.Context) {
	user := &model.User{}
	if c.BindJSON(user) != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.CreateUser", "input params", "invalid json body"))
		return
	}
	if user, err := biz.CreateUser(user); err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   200,
			"result": "ok",
			"user":   user,
		})
	}
}
