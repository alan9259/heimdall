package handler

import (
	platform "heimdall/internal/platform"
	middleware "heimdall/internal/router/middleware"

	"github.com/labstack/echo"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(platform.JWTSecret, h.revokedTokenStore)
	guestUsers := v1.Group("/accounts")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)
	guestUsers.POST("/logout", h.Logout, jwtMiddleware)
	guestUsers.POST("/change", h.Change, jwtMiddleware)
	guestUsers.POST("/verify", h.Verify)
	guestUsers.POST("/forgot", h.ForgotPassword)
	guestUsers.POST("/validate-pin", h.ValidatePin)

	profiles := v1.Group("/profiles", jwtMiddleware)
	profiles.GET("/current", h.GetCurrentAccount)
	profiles.PUT("", h.UpdateAccount)

	locations := v1.Group("/locations", middleware.JWTWithConfig(
		middleware.JWTConfig{
			Skipper: func(c echo.Context) bool {
				if c.Request().Method == "GET" && c.Path() != "/api/locations" {
					return true
				}
				return false
			},
			SigningKey: platform.JWTSecret,
		},
		h.revokedTokenStore,
	))

	locations.POST("", h.SignUp)

	/*articles.POST("", h.CreateArticle)
	articles.GET("/feed", h.Feed)
	articles.PUT("/:slug", h.UpdateArticle)
	articles.DELETE("/:slug", h.DeleteArticle)
	articles.POST("/:slug/comments", h.AddComment)
	articles.DELETE("/:slug/comments/:id", h.DeleteComment)
	articles.POST("/:slug/favorite", h.Favorite)
	articles.DELETE("/:slug/favorite", h.Unfavorite)
	articles.GET("", h.Articles)
	articles.GET("/:slug", h.GetArticle)
	articles.GET("/:slug/comments", h.GetComments)

	tags := v1.Group("/tags")
	tags.GET("", h.Tags) */
}
