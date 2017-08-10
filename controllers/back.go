package controllers

import (
	"github.com/EtienneR/gin_session/db"
	"github.com/EtienneR/gin_session/helpers"

	"github.com/gin-gonic/gin"
)

func AdminHandler(c *gin.Context) {
	var users []db.Users

	dbmap := db.InitDb()
	defer dbmap.Close()
	dbmap.Find(&users)

	// Récupération du nom d'utilisateur pour le templating
	userName := helpers.GetUserName(c)

	c.HTML(200, "view_admin.html", gin.H{
		"userName":    userName,
		"currentPage": "admin",
		"users":       users,
		"success":     helpers.GetFlashCookie(c, "success"),
	})
}
