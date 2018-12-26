package biz

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/KenmyZhang/smart-edu-server/common/config"
	"github.com/KenmyZhang/smart-edu-server/common/util"
	"github.com/KenmyZhang/smart-edu-server/log"
	"github.com/KenmyZhang/smart-edu-server/model"
	"strings"
	"time"
	"github.com/KenmyZhang/golang-lib/middleware"
)

const (
	SendMessageUrl = `https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=`
	AccessTokenUrl = `https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential`
	MaxRequestTry  = 5
)

func SendMessage(msg *model.SendMsg) *util.Err {
	token, err := GetAccessToken()
	if err != nil {
		return err
	}

	req, _ := http.NewRequest("POST", SendMessageUrl+token, strings.NewReader(msg.ToJson()))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	log.Info(fmt.Sprintf("Sending message: %+v", msg))
	var resp *http.Response
	start := time.Now()
	for i := 0; i < MaxRequestTry; i++ {
		var err error
		resp, err = util.NewClient().Do(req)
		if err != nil {
			if i < MaxRequestTry {
				continue
			} else {
				end := time.Now()
				latency := end.Sub(start)
				log.Error(fmt.Sprintf("POST %s | %13v | timeout", SendMessageUrl, latency))
				middleware.ErrCounter.WithLabelValues("POST", SendMessageUrl).Inc()
				log.Error("Sending message failed, " + err.Error())
				return util.NewInternalServerError("biz.SendMessage", err.Error())
			}
		}
	}

	end := time.Now()
	latency := end.Sub(start)
	log.Error(fmt.Sprintf("POST %s | %13v ", SendMessageUrl, latency))

	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Error("ReadAll message failed, " + err.Error())
		return util.NewInternalServerError("biz.SendMessage", err.Error())
	} else {
		log.Info("Sending message successfully," + string(data))
	}
	return nil
}

func ActionAfterRecvMsg(recvMsg *model.RecvMsg) *util.Err {
	sendMsg := &model.SendMsg{}
	sendMsg.ToUser = recvMsg.FromUserName
	sendMsg.MsgType = "text"
	sendMsg.Text.Content = "Hello World"
	var err *util.Err
	if err = SendMessage(sendMsg); err == nil {
		return nil
	}
	return err
}

func GetAccessToken() (string, *util.Err) {
	req, _ := http.NewRequest("GET", AccessTokenUrl+`&appid=`+config.Cfg.WexinSettings.AppId+`&secret=`+config.Cfg.WexinSettings.AppSecret, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if resp, err := util.NewClient().Do(req); err != nil {
		log.Error("Get access token failed" + err.Error())
		return "", util.NewInternalServerError("biz.GetAccessToken", err.Error())
	} else {
		if data, err := ioutil.ReadAll(resp.Body); err != nil {
			log.Error("ReadAll message failed, " + err.Error())
			return "", util.NewInternalServerError("biz.GetAccessToken", err.Error())
		} else {
			result := &model.AccessTokenResp{}
			err = json.Unmarshal([]byte(data), result)
			if err != nil {
				return "", util.NewInternalServerError("biz.GetAccessToken", err.Error())
			}
			log.Info("Get access token message successfully," + result.AccessToken)
			return result.AccessToken, nil
		}
	}
}
