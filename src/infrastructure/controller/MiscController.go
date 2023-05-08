package controller

import (
	"strconv"
	"net/http"
	"go-ascii/src/commons/log"
	"github.com/gin-gonic/gin"
	"go-ascii/src/commons/constants/log-categories"
)

const Family = "API"

func logRequest(c *gin.Context, family string, status int) {
	category := log_categories.INFO
	if status != http.StatusOK {
		category = log_categories.WARNING
	}

	log.LogFam(category, family, ""+c.Request.Method+" petition to end point '"+c.Request.URL.Path+"'. Status: "+ strconv.Itoa(status))
}