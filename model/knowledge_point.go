package model

import (
	"fmt"
	log "github.com/KenmyZhang/golang-lib/zaplogger"
	"github.com/KenmyZhang/smart-edu-server/common/util"
)

type KnowledgePoint struct {
	Id          string `json:"id"`
	Label       string `json:"label"`
	ParentId    string `json:"parent_id"`
	CreatedTime int64  `json:"created_time"    bson:"created_time"`
	UpdatedTime int64  `json:"created_time"    bson:"updated_time"`
	DeletedTime int64  `json:"deleted_time"    bson:"deleted_time"`
}

type MultiLevelKnowledgePoint struct {
	Id       string                     `json:"id"`
	Label    string                     `json:"label"`
	Children KnowledgePointList `json:"children"`
}

type KnowledgePointList []MultiLevelKnowledgePoint

func (c *KnowledgePoint) PreSave() {
	c.Id = util.NewId()
	c.CreatedTime = util.GetMillis()
	c.UpdatedTime = util.GetMillis()
	c.DeletedTime = 0
}

func CreateKnowlegePoint(k *KnowledgePoint) (*KnowledgePoint, *util.Err) {
	k.PreSave()
	if result := sqlSupplier.GetMaster().Model(&KnowledgePoint{}).Create(k); result.Error != nil {
		log.Error(fmt.Sprintf("CreateKnowlegePoint fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.CreateKnowlegePoint", result.Error.Error())
	}
	return k, nil
}

func GetKnowledgePoint(id string) (*KnowledgePoint, *util.Err) {
	k := &KnowledgePoint{}
	if result := sqlSupplier.GetReplica().Where("id = ?", id).First(k); result.Error != nil {
		log.Error(fmt.Sprintf("GetKnowledgePoint fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.GetKnowledgePoint", result.Error.Error())
	}

	return k, nil
}

func GetChildKnowledgePoints(id string) (*KnowledgePointList, *util.Err) {
	list := &KnowledgePointList{}
	if result := sqlSupplier.GetMaster().Where("parent_id = ?", id).Order("created_time desc").Find(list); result.Error != nil {
		log.Error(fmt.Sprintf("GetChildKnowledgePoints fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.GetChildKnowledgePoints", result.Error.Error())
	}

	return list, nil
}
