package biz

import (
	"io/ioutil"
	"github.com/KenmyZhang/smart-edu-server/common/config"
	"github.com/KenmyZhang/smart-edu-server/common/util"
)

func GetArticle(articleId string) ([]byte, *util.Err) {
	path := config.Cfg.FilePath + articleId
	if data, err := ReadFile(path); err != nil {
		return nil, err
	} else {
		return data, nil
	}

}
func ReadFile(path string) ([]byte, *util.Err) {
	if f, err := ioutil.ReadFile(path); err != nil {
		return nil, util.NewInternalServerError("biz.ReadFile", err.Error())
	} else {
		return f, nil
	}
}
