package model

import (
	"fmt"
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/log"
)

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Role        string `json:"role"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	UnionId     string `json:"union_id"`
	Description string `json:"description"`
	CreatedTime int64  `json:"created_time"`
	DeletedTime int64  `json:"deleted_time"`
	UpdatedTime int64  `json:"updated_time"`
}

func (c *User) PreSave() {
	c.Id = util.NewId()
	c.CreatedTime = util.GetMillis()
	c.UpdatedTime = util.GetMillis()
}

func (c *User) Update() {
	c.UpdatedTime = util.GetMillis()
}

func CreateUser(u *User) (*User, *util.Err) {
	u.PreSave()
	if result := sqlSupplier.GetMaster().Model(&User{}).Create(u); result.Error != nil {
		log.Error(fmt.Sprintf("create user fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.CreateUser", result.Error.Error())
	}
	return u, nil
}

func DeleteUser(UserID string) *util.Err {
	if result := sqlSupplier.GetMaster().Where("id = ?", UserID).Delete(&User{}); result.Error != nil {
		log.Error(fmt.Sprintf("delete user, %s", result.Error.Error()))
		return util.NewInternalServerError("model.DeleteUser", result.Error.Error())
	}
	return nil
}

func GetUser(userID string) (*User, *util.Err) {
	user := &User{}
	if result := sqlSupplier.GetReplica().Where("id = ?", userID).First(user); result.Error != nil {
		log.Error(fmt.Sprintf("get user fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.GetUser", result.Error.Error())
	}

	return user, nil
}

func GetUsers(userIDs []string) ([]*User, *util.Err) {
	var users []*User
	if result := sqlSupplier.GetMaster().Where("id in (?)", userIDs).Order("created_at desc").Find(&users); result.Error != nil {
		log.Error(fmt.Sprintf("get users fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.GetUsers", result.Error.Error())
	}

	return users, nil
}

func GetAllUsers(offset, limit int64) ([]*User, *util.Err) {
	var users []*User
	if result := sqlSupplier.GetMaster().Offset(offset).Limit(limit).Find(&users); result.Error != nil {
		log.Error(fmt.Sprintf("get all users fail, %s", result.Error.Error()))
		return nil, util.NewInternalServerError("model.GetUsers", result.Error.Error())
	}
	return users, nil
}
