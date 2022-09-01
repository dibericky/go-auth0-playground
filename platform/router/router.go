package router

import (
	"encoding/gob"
	"go-auth0-playground/platform/authenticator"
	"go-auth0-playground/platform/middleware"
	"go-auth0-playground/web/app/callback"
	"go-auth0-playground/web/app/home"
	"go-auth0-playground/web/app/login"
	"go-auth0-playground/web/app/logout"
	"go-auth0-playground/web/app/user"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// New registers the routes and returns the router.
func New(auth *authenticator.Authenticator) *gin.Engine {
	router := gin.Default()

	// To store custom types in our cookies,
	// we must first register them using gob.Register
	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.Static("/public", "web/static")
	router.LoadHTMLGlob("web/template/*")

	router.GET("/", home.Handler)
	router.GET("/login", login.Handler(auth))
	router.GET("/callback", callback.Handler(auth))
	router.GET("/user", middleware.IsAuthenticated, user.Handler)
	router.GET("/logout", logout.Handler)

	return router
}
