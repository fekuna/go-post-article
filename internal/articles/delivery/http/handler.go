package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fekuna/go-post-article/config"
	"github.com/fekuna/go-post-article/internal/articles"
	"github.com/fekuna/go-post-article/internal/models"
	"github.com/fekuna/go-post-article/pkg/httpResponse"
	"github.com/fekuna/go-post-article/pkg/logger"
	"github.com/fekuna/go-post-article/pkg/utils"
	"github.com/labstack/echo/v4"
)

// articles handlers
type articlesHandlers struct {
	cfg        *config.Config
	logger     logger.Logger
	articlesUC articles.UseCase
}

func NewArticlesHandlers(cfg *config.Config, logger logger.Logger, articlesUC articles.UseCase) articles.Handlers {
	return &articlesHandlers{
		cfg:        cfg,
		logger:     logger,
		articlesUC: articlesUC,
	}
}

func (h *articlesHandlers) AddArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		article := &models.Article{}
		if err := utils.ReadRequest(c, article); err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		lastInsertedId, err := h.articlesUC.AddArticle(ctx, article)
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		article.Id = int(*lastInsertedId)
		return httpResponse.Success(c, http.StatusCreated, article, "Success to insert article")
	}
}

func (h *articlesHandlers) FindArticleById() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		articleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		foundArticle, err := h.articlesUC.FindArticleById(ctx, articleId)
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		return httpResponse.Success(c, http.StatusOK, foundArticle, "Success to get article")
	}
}

func (h *articlesHandlers) UpdateArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		articleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		articleData := &models.Article{Id: articleId}

		if err := utils.ReadRequest(c, articleData); err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		_, err = h.articlesUC.UpdateArticle(ctx, articleData)
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		return httpResponse.Success(c, http.StatusCreated, articleData, "Success to update article")
	}
}

func (h *articlesHandlers) DeleteArticle() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		articleId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		err = h.articlesUC.DeleteArticle(ctx, articleId)
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		return httpResponse.Success(c, http.StatusOK, nil, "Success to delete article")
	}
}

func (h *articlesHandlers) GetArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()

		paginationQuery, err := utils.GetPaginationFromCtx(c)
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		articles, err := h.articlesUC.GetArticles(ctx, paginationQuery)
		if err != nil {
			return httpResponse.ErrorWithLog(c, h.logger, err)
		}

		return httpResponse.SuccessPagination(c, http.StatusOK, articles.Articles, articles.Meta, "Success to get list articles")
	}
}
