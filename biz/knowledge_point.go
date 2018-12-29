package biz

import (
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/model"
)

func CreateKnowlegePoint(k *model.KnowledgePoint) (*model.KnowledgePoint, *util.Err) {
	if knowledgePoint, err := model.CreateKnowlegePoint(k); err != nil {
		return nil, err
	} else {
		return knowledgePoint, nil
	}
}

func GetKnowledgePoint(id string) (*model.KnowledgePoint, *util.Err) {
	if knowledgePoint, err := model.GetKnowledgePoint(id); err != nil {
		return nil, err
	} else {
		return knowledgePoint, nil
	}
}

func GetChildKnowledgePoints(id string) (*model.KnowledgePointList, *util.Err) {
	if list, err := model.GetChildKnowledgePoints(id); err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

func GetChildKnowledgePointAndChild(id string) (*model.MultiLevelKnowledgePoint, *util.Err) {
	multiLevelKnowledgePoint := &model.MultiLevelKnowledgePoint{}

	if knowledgePoint, err := model.GetKnowledgePoint(id); err != nil {
		return nil, err
	} else {
		multiLevelKnowledgePoint.Id = knowledgePoint.Id
		multiLevelKnowledgePoint.Label = knowledgePoint.Label
	}

	if list, err := model.GetChildKnowledgePoints(id); err != nil {
		return nil, err
	} else {
		multiLevelKnowledgePoint.Children = append(multiLevelKnowledgePoint.Children, *list...)
		return multiLevelKnowledgePoint, nil
	}
}
