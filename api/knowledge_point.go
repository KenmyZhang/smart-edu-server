package api

import (
	"github.com/KenmyZhang/smart-edu-server/biz"
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateKnowlegePoint(c *gin.Context) {
	knowledgePoint := &model.KnowledgePoint{}
	if c.BindJSON(knowledgePoint) != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.CreateKnowlegePoint", "input params", "invalid json body"))
		return
	}
	if rst, err := biz.CreateKnowlegePoint(knowledgePoint); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":            200,
			"knowledge_point": rst,
		})
	}
}

func GetKnowlegePointById(c *gin.Context) {
	knowledgePoint := &model.KnowledgePoint{}
	if c.BindJSON(knowledgePoint) != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.GetKnowlegePointById", "input params", "invalid json body"))
		return
	}
	if rst, err := biz.GetKnowledgePoint(knowledgePoint.Id); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":            200,
			"knowledge_point": rst,
		})
	}
}

func GetChildKnowledgePoints(c *gin.Context) {
	knowledgePoint := &model.KnowledgePoint{}
	if c.BindJSON(knowledgePoint) != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.GetChildKnowledgePoints", "input params", "invalid json body"))
		return
	}
	if list, err := biz.GetChildKnowledgePoints(knowledgePoint.Id); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"list": list,
		})
	}
}

func GetChildKnowledgePointAndChild(c *gin.Context) {
	knowledgePoint := &model.KnowledgePoint{}
	if c.BindJSON(knowledgePoint) != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.GetChildKnowledgePoints", "input params", "invalid json body"))
		return
	}
	if list, err := biz.GetChildKnowledgePointAndChild(knowledgePoint.Id); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"list": list,
		})
	}
}
