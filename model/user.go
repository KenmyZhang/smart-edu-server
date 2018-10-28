package model

import (
	"fmt"
	"smart-edu-server/common/util"
	"smart-edu-server/log"
)

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	Location  string `json:"location"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	UnionId   string `json:"union_id"`
	CreatedAt int64  `json:"created_at"`
	DeletedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func (c *User) PreSave() {
	c.CreatedAt = util.GetMillis()
	c.UpdatedAt = util.GetMillis()
}

func (c *User) Update() {
	c.UpdatedAt = util.GetMillis()
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
