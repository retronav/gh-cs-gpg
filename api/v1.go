package handler

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.karawale.in/gh-cs-gpg/middleware"
)

func V1Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	gpgGroup := v1.Group("/gpg")
	gpgGroup.Use(middleware.GitHubMiddleware)

	gpgGroup.GET("/priv.gpg", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, os.Getenv("PRIV_GPG_KEY"))
	})
	gpgGroup.GET("/pub.gpg", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, os.Getenv("PUB_GPG_KEY"))
	})

	sshGroup := v1.Group("/ssh")
	sshGroup.Use(middleware.GitHubMiddleware)

	sshGroup.GET("/gh_codespaces.id_ed25519", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, os.Getenv("PRIV_SSH_KEY"))
	})
	sshGroup.GET("/gh_codespaces.id_ed25519.pub", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, os.Getenv("PUB_SSH_KEY"))
	})

	router.ServeHTTP(w, r)
}
