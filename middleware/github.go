package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
)

func GitHubMiddleware(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.String(
			http.StatusUnauthorized,
			"Don't try to break in without an Authorization header.",
		)
		ctx.Abort()
		return
	}
	tokenParts := strings.Split(header, "Bearer ")
	if len(tokenParts) != 2 {
		ctx.String(
			http.StatusUnauthorized,
			"Don't try to break in without a proper Authorization header.",
		)
		ctx.Abort()
		return
	}
	login := os.Getenv("GH_LOGIN")
	if login == "" {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	oauthClient := oauth2.NewClient(
		ctx, oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: tokenParts[1]},
		))
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if user.GetLogin() != login {
		ctx.String(
			http.StatusUnauthorized,
			"GitHub token OK but you're not the user I wanted.",
		)
		ctx.Abort()
		return
	}
	ctx.Next()
}
