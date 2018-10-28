package api

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"smart-edu-server/biz"
	"smart-edu-server/common/util"
	"smart-edu-server/model"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func SendMsg(c *gin.Context) {
	revmsg := &model.RecvMsg{}
	if err := c.Bind(revmsg); err != nil {
		c.JSON(http.StatusOK, util.NewInvalidParamError(util.InvalidParam, "api.RecvMsg", "input params", "invalid json body, "+err.Error()))
		return
	}

	if err := biz.ActionAfterRecvMsg(revmsg); err != nil {
		c.JSON(http.StatusOK, err)
		return
	} else {
		c.JSON(http.StatusOK, "success")
	}
}

func CheckServer(c *gin.Context) {
	echostr := c.Query("echostr")
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	token := model.MessageToken

	arr := []string{timestamp, nonce, token}

	sort.Strings(arr)

	arrStr := strings.Join(arr, "")

	h := sha1.New()
	io.WriteString(h, arrStr)
	r := h.Sum(nil)

	if signature == hex.EncodeToString(r[:]) {
		num, _ := strconv.Atoi(echostr)
		c.JSON(http.StatusOK, num)
	} else {
		c.JSON(http.StatusOK, false)
	}
}
