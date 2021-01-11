package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/linkc0829/go-ics/internal/db/mongodb"
	"github.com/linkc0829/go-ics/internal/handlers"
	"github.com/linkc0829/go-ics/internal/handlers/auth"
	"github.com/linkc0829/go-ics/internal/handlers/secret"
	"github.com/linkc0829/go-ics/pkg/utils"
)

func Auth(cfg *utils.ServerConfig, r *gin.Engine, db *mongodb.MongoDB) {

	// OAuth handlers
	g := r.Group(cfg.VersioningEndpoint("/auth"))
	g.GET("/:provider", auth.Begin())
	g.GET("/:provider/callback", auth.CallBack(cfg, db))

	// ics secrets handler
	g.POST("/:provider/signup", secret.SignupHandler(cfg, db))
	g.POST("/:provider/login", secret.LoginHandler(cfg, db))
	g.GET("/:provider/refresh_token", secret.RefreshTokenHandler(cfg, db))

	g.GET("/:provider/logout", handlers.LogoutHandler())

}
