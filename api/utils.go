package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"smart-edu-server/common/config"
	"smart-edu-server/common/util"
)

func GetConfig(c *gin.Context) {
	cfg := config.GetConfig()
	c.JSON(http.StatusOK, gin.H{
		"result": cfg,
		"code":   "200 OK",
	})
}

func GetVersionDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"result": gin.H{"CurrentVersion": util.CurrentVersion, "BuildDate": util.BuildDate, "BuildHash": util.BuildHash},
		"code":   "200 OK",
	})
}
