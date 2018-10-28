package util

import (
	"bytes"
	"encoding/base32"
	"encoding/gob"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/pborman/uuid"
)

var versions = []string{
	"1.0",
}

var CLOUD_LIST_PRE = "xlppc.smart-edu-server.api"

var CurrentVersion string = versions[0]
var BuildDate string
var BuildHash string

const (
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"
)

func init() {
	gob.Register(map[string]interface{}{})
}

// GetMillis is a convience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeBytes(input []byte, thing interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(input))
	return dec.Decode(thing)
}

const (
	connectTimeout        = 6 * time.Second
	requestTimeout        = 6 * time.Second
	responseHeaderTimeout = 6 * time.Second
)

func NewClient() *http.Client {
	dialContext := (&net.Dialer{
		Timeout:   connectTimeout,
		KeepAlive: 30 * time.Second,
	}).DialContext

	client := &http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ResponseHeaderTimeout: responseHeaderTimeout,
		},
		Timeout: requestTimeout,
	}
	return client
}

func checkFileExit(filename string) bool {
	var exit = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exit = false
	}
	return exit
}

func GetFile(filename string) (*os.File, error) {
	if checkFileExit(filename) {
		return os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0666)
	} else {
		index := strings.LastIndex(filename, "/")
		// 创建文件夹
		err := os.MkdirAll(filename[:index], os.ModePerm)
		if err != nil {
			return nil, err
		}
		return os.Create(filename)
	}

}

const (
	InvalidParam             = 4000
	InvalidResources         = 4001
	InvalidResourceUrl       = 4002
	NoPerToGetResources      = 4003 //没有权限访问资源
	NoPerToChangeResource    = 4004
	NoPerToAddResources      = 4005
	InvalidResourcename      = 4006
	InvalidSessionId         = 4007
	NoPerAddCollectionMember = 4008
	PermissionError          = 4009
	InternalServerError      = 5000
)

type Err struct {
	Code   int    `json:"code"`
	Result string `json:"result"`
}

func (e *Err) Error() string {
	if e == nil {
		return ""
	}
	return "code:" + strconv.Itoa(e.Code) + ", result:" + e.Result
}

func NewInvalidParamError(code int, where, parameter, details string) *Err {
	var message string
	if details != "" {
		message = ", details:" + details + ", where:" + where
	} else {
		message = ", where:" + where
	}
	return &Err{Code: code, Result: "Invalid " + parameter + " patameter" + message}
}

func NewInternalServerError(where, details string) *Err {
	var message string
	if details != "" {
		message = "details:" + details + ", where:" + where
	} else {
		message = "where:" + where
	}
	return &Err{Code: InternalServerError, Result: "Internal Server Error," + message}
}

var encoding = base32.NewEncoding("abnerfg8ejkncpqxot1uwisza345h769")

// NewId is a globally unique identifier.  It is a [A-Z0-9] string 26
// characters long.  It is a UUID version 4 Guid that is zbased32 encoded
// with the padding stripped off.
func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}
