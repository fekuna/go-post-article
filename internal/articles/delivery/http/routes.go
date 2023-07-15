package http

import (
	"github.com/fekuna/go-post-article/internal/articles"
	"github.com/labstack/echo/v4"
)

func MapArticlesRoutes(authGroup *echo.Group, h articles.Handlers) {
	authGroup.POST("", h.AddArticle())
	authGroup.GET("", h.GetArticles())
	authGroup.GET("/:id", h.FindArticleById())
	authGroup.PUT("/:id", h.UpdateArticle())
	authGroup.DELETE("/:id", h.DeleteArticle())
}
