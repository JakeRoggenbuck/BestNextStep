package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "HomePage", nil)
}

func apiRootPage(c *gin.Context) {
	c.HTML(http.StatusOK, "ApiRootPage", nil)
}
