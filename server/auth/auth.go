package auth

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/stephane-nguyen/FamilyTree/server/config"
)

const (
	maxAge = 86400 * 30 // 30 days
	isProd = false
)

func OAuth2() {
	env := config.LoadEnv()
	googleClientId := env.GOOGLE_CLIENT_ID
	googleClientSecret := env.GOOGLE_CLIENT_SECRET
	googleCallbackUrl := env.GOOGLE_CALLBACK_URL

	googleProvider := google.New(googleClientId, googleClientSecret, googleCallbackUrl)
	goth.UseProviders(googleProvider)
}