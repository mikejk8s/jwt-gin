package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginPost(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
