package controllers

import (
	"github.com/EtienneR/gin_session/db"
	"github.com/EtienneR/gin_session/helpers"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func LoginHandlerForm(c *gin.Context) {
	// Récupération du nom d'utilisateur pour le templating
	userName := helpers.GetUserName(c)

	c.HTML(200, "view_form.html", gin.H{
		"userName":       userName,
		"currentPage":    "login",
		"warning":        helpers.GetFlashCookie(c, "warning"),
		csrf.TemplateTag: csrf.TemplateField(c.Request),
	})
}

func LoginHandler(c *gin.Context) {
	var user db.Users

	// Binding du formulaire
	if c.Bind(&user) == nil {
		// Récupération du mdp en clair
		clearPassword := user.Password

		// Connexion au fichier SQLite
		dbmap := db.InitDb()
		defer dbmap.Close()

		if err := dbmap.Where("name = ?", user.Name).First(&user).Error; err == nil {

			// Vérification du mdp
			if helpers.CheckPasswordHash(clearPassword, user.Password) {
				// Création du cookie de session
				helpers.SetSession(user.Name, c)
				helpers.SetFlashCookie(c, "success", "Bienvenue "+user.Name)
				c.Redirect(302, "/admin")
			} else {
				// MDP incorrect
				helpers.SetFlashCookie(c, "warning", "Mot de passe incorrect")
				c.Redirect(302, "/login")
			}

		} else {
			// Nom d'utilisateur incorrect
			helpers.SetFlashCookie(c, "warning", "Nom d'utilisateur incorrect")
			c.Redirect(302, "/login")
		}

	} else {
		// Champs non remplis
		helpers.SetFlashCookie(c, "warning", "Champs non remplis")
		c.Redirect(302, "/login")
	}
}

func LogoutHandler(c *gin.Context) {
	helpers.ClearSession(c)
	helpers.SetFlashCookie(c, "success", "Vous êtes désormais déconnecté(e)")
	c.Redirect(302, "/")
}

func SignUpHandlerForm(c *gin.Context) {
	// Récupération du nom d'utilisateur pour le templating
	userName := helpers.GetUserName(c)

	c.HTML(200, "view_form.html", gin.H{
		"userName":       userName,
		"currentPage":    "signUp",
		"warning":        helpers.GetFlashCookie(c, "warning"),
		csrf.TemplateTag: csrf.TemplateField(c.Request),
	})
}

func SignUpHandler(c *gin.Context) {
	var user db.Users

	if c.Bind(&user) == nil {
		// Mot de passe haché
		hash, _ := helpers.HashPassword(user.Password)
		user.Password = hash

		// Connexion à SQLite
		dbmap := db.InitDb()
		defer dbmap.Close()

		if err := dbmap.Where("name = ?", user.Name).First(&user).Error; err == nil {
			helpers.SetFlashCookie(c, "warning", "Inscription refusée, le nom d'utitlisateur "+user.Name+" existe déjà")
			c.Redirect(302, "/signup")
		} else {
			// Création de l'utilisateur
			dbmap.Create(&user)
			// Création du cookie de session
			helpers.SetSession(user.Name, c)
			c.Redirect(302, "/admin")
		}

	} else {
		helpers.SetFlashCookie(c, "warning", "Champs non remplis")
		c.Redirect(302, "/signup")
	}
}
