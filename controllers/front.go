package controllers

import (
	"github.com/EtienneR/gin_session/helpers"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	userName := helpers.GetUserName(c)

	c.HTML(200, "view_index.html", gin.H{
		"userName":    userName,
		"currentPage": "index",
		"success":     helpers.GetFlashCookie(c, "success"),
	})
}
