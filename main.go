package main

import (
	"net/http"

	"github.com/EtienneR/gin_session/controllers"
	"github.com/EtienneR/gin_session/helpers"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/csrf"
)

func Private() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Request.Cookie(helpers.CookieSessionName)

		if err != nil {
			c.Redirect(302, "/login")
			c.Abort() // Très important (sinon pas protégée avec curl -i http://localhost:3000/admin)
		}
	}
}

func main() {
	// Initialisation du routeur
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	// Page d'accueil
	r.GET("/", controllers.IndexHandler)

	// Page admin (privée)
	admin := r.Group("/admin") // Nom du groupe "admin" dont l'URI est "/admin"
	admin.Use(Private())       // Utilisation du middleware Private()
	{
		// Page admin (privée)
		admin.GET("", controllers.AdminHandler)
		// Autres routes dans "/admin"
	}

	// Page de connexion
	r.GET("/login", controllers.LoginHandlerForm)
	r.POST("/login", controllers.LoginHandler)

	// Page d'inscription
	r.GET("/signup", controllers.SignUpHandlerForm)
	r.POST("/signup", controllers.SignUpHandler)

	// Page de déconnexion
	r.GET("/logout", controllers.LogoutHandler)

	// Port du serveur
	//r.Run(":3000")
	http.ListenAndServe(":3000",
		csrf.Protect([]byte("32-byte-long-auth-key"), csrf.Secure(false))(r))
}
