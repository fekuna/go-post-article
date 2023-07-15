package articles

import "github.com/labstack/echo/v4"

// Articles HTTP Handlers interface
type Handlers interface {
	AddArticle() echo.HandlerFunc
	FindArticleById() echo.HandlerFunc
	UpdateArticle() echo.HandlerFunc
	DeleteArticle() echo.HandlerFunc
	GetArticles() echo.HandlerFunc
}
