package model

import (
	"encoding/json"
)

type SendMsg struct {
	ToUser  string  `json:"touser"`  //普通用户(openid)
	MsgType string  `json:"msgtype"` //消息类型，文本为text，图文链接为link
	Text    MsgText `json:"text"`
}

const MessageToken = "7m0IjnadDrUaXP6rJoJm0Qjwes"

type MsgText struct {
	Content string `json:"content"`
}

func (o *SendMsg) ToJson() string {
	b, _ := json.Marshal(o)
	return string(b)
}

type RecvMsg struct {
	MsgId        int64  `json:"MsgId" xml:"MsgId"`
	ToUsername   string `json:"ToUserName" xml:"ToUserName"`
	FromUserName string `json:"FromUserName" xml:"FromUserName"`
	CreateTime   int64  `json:"CreateTime" xml:"CreateTime"`
	MsgType      string `json:"MsgType" xml:"MsgType"`
	Event        string `json:"Event" xml:"Event"`
	SessionFrom  string `json:"SessionFrom" xml:"SessionFrom"`
	Title        string `json:"title" xml":"title`
	AppId        string `json:"appid" xml:"appid"`
	PagePath     string `json:"PagePath" xml:"PagePath"`
	ThumbUrl     string `json:"thumbUrl" xml:"thumbUrl"`
	ThumbMediaId string `json:"thumbMediaId" xml:"thumbMediaId"`
	PicUrl       string `json:"PicUrl" xml:"PicUrl"`
	Mediaid      string `json:"MediaId" xml:"MediaId"`
	Content      string `json:"content" xml:"content"`
}

type AccessTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

const (
	WXBizMsgCrypt_OK                      = 0
	WXBizMsgCrypt_ValidateSignature_Error = -40001
	WXBizMsgCrypt_ParseXml_Error          = -40002
	WXBizMsgCrypt_ComputeSignature_Error  = -40003
	WXBizMsgCrypt_IllegalAesKey           = -40004
	WXBizMsgCrypt_ValidateAppid_Error     = -40005
	WXBizMsgCrypt_EncryptAES_Error        = -40006
	WXBizMsgCrypt_DecryptAES_Error        = -40007
	WXBizMsgCrypt_IllegalBuffer           = -40008
	WXBizMsgCrypt_EncodeBase64_Error      = -40009
	WXBizMsgCrypt_DecodeBase64_Error      = -40010
	WXBizMsgCrypt_GenReturnXml_Error      = -40011
)
