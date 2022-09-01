package login

import (
	"crypto/rand"
	"encoding/base64"
	"go-auth0-playground/platform/authenticator"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our login.
func Handler(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// doing this you can't handle redict on frontend manually
		// ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))

		ctx.JSON(http.StatusOK, gin.H{
			"url": auth.AuthCodeURL(state),
		})
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}