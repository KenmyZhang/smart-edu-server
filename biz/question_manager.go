package biz

import (
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/model"
)

func CreateQuestion(question *model.Question) *util.Err {
	if err := model.SaveQuestion(question); err != nil {
		return err
	} else {
		return nil
	}
}

func GetQuestions(offset, limit int) (*model.QuestionList, *util.Err) {
	if list, err := model.GetQuestions(offset, limit); err != nil {
		return nil, err
	} else {
		return list, nil
	}
}
