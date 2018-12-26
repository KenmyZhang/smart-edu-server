package model

import (
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"gopkg.in/mgo.v2/bson"
)

type Question struct {
	Id             string   `json:"id" 				bson:"_id"`
	Question       string   `json:"question" 		bson:"question"`
	KnowlegePoints []string `json:"knowlege_point" 	bson:"knowledge_points"`
	Answer         string   `json:"answer" 			bson:"answer"`
	Option         []string `json:"option" 			bson:"option"`
	Type           string   `json:"type"		    bson:"type"`
	Analysis       string   `json:"analysis"        bson:"analysis"`
	CreatedTime    int64    `json:"created_time"    bson:"created_time"`
	UpdatedTime    int64    `json:"created_time"    bson:"updated_time"`
}

type QuestionList []Question

func (c *Question) PreSave() {
	c.Id = util.NewId()
	c.CreatedTime = util.GetMillis()
	c.UpdatedTime = util.GetMillis()
}

func SaveQuestion(question *Question) *util.Err {
	question.PreSave()
	err := QuestionCollection.Insert(question)
	if err != nil {
		return util.NewInternalServerError("model.SaveQuestion", err.Error())
	} else {
		return nil
	}
}

func GetQuestions(offset, limit int) (*QuestionList, *util.Err) {
	questionList := &QuestionList{}
	if err := QuestionCollection.Find(&bson.M{}).Sort("-CreatedTime").Limit(100).All(questionList); err != nil {
		return nil, util.NewInternalServerError("model.GetQuestions", err.Error())
	} else {
		return questionList, nil
	}
}
