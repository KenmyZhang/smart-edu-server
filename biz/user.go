package biz

import (
	"smart-edu-server/common/util"
	"smart-edu-server/model"
)

func GetUser(userID string) (*model.User, *util.Err) {
	if user, err := model.GetUser(userID); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func CreateUser(user *model.User) (*model.User, *util.Err) {
	if rst, err := model.CreateUser(user); err != nil {
		return nil, err
	} else {
		return rst, nil
	}
}
