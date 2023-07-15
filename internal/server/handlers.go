package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	articlesHttp "github.com/fekuna/go-post-article/internal/articles/delivery/http"
	articlesRepo "github.com/fekuna/go-post-article/internal/articles/repository"
	articlesUC "github.com/fekuna/go-post-article/internal/articles/usecase"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repository
	articlesRepo := articlesRepo.NewArticleRepository(s.db)

	// Init useCase
	articlesUC := articlesUC.NewArticlesUseCase(s.cfg, s.logger, articlesRepo)

	// Init handlers
	articlesHandlers := articlesHttp.NewArticlesHandlers(s.cfg, s.logger, articlesUC)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))

	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	// v1 := e.Group("/api/v1")

	articlesGroup := e.Group("/article")

	articlesHttp.MapArticlesRoutes(articlesGroup, articlesHandlers)

	return nil
}
