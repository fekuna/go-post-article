package usecase

import (
	"context"

	"github.com/fekuna/go-post-article/config"
	"github.com/fekuna/go-post-article/internal/articles"
	"github.com/fekuna/go-post-article/internal/models"
	"github.com/fekuna/go-post-article/pkg/logger"
	"github.com/fekuna/go-post-article/pkg/utils"
)

// Articles UseCase
type articlesUC struct {
	cfg          *config.Config
	logger       logger.Logger
	articlesRepo articles.Repository
}

// Articles UseCase constructor
func NewArticlesUseCase(cfg *config.Config, logger logger.Logger, articlesRepo articles.Repository) articles.UseCase {
	return &articlesUC{
		cfg:          cfg,
		logger:       logger,
		articlesRepo: articlesRepo,
	}
}

func (u *articlesUC) AddArticle(ctx context.Context, article *models.Article) (*int64, error) {
	insertedId, err := u.articlesRepo.AddArticle(ctx, article)
	if err != nil {
		return nil, err
	}

	return insertedId, nil
}

func (u *articlesUC) FindArticleById(ctx context.Context, articleId int) (*models.Article, error) {
	foundArticle, err := u.articlesRepo.FindArticleById(ctx, articleId)
	if err != nil {
		return nil, err
	}

	return foundArticle, nil
}

func (u *articlesUC) UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {
	_, err := u.articlesRepo.FindArticleById(ctx, article.Id)
	if err != nil {
		return nil, err
	}

	updatedArticle, err := u.articlesRepo.UpdateArticle(ctx, article)
	if err != nil {
		return nil, err
	}

	return updatedArticle, nil
}

func (u *articlesUC) DeleteArticle(ctx context.Context, articleId int) error {
	_, err := u.articlesRepo.FindArticleById(ctx, articleId)
	if err != nil {
		return err
	}

	err = u.articlesRepo.DeleteArticleById(ctx, articleId)
	if err != nil {
		return err
	}

	return nil
}

func (u *articlesUC) GetArticles(ctx context.Context, pq *utils.PaginationQuery) (*models.ArticleList, error) {

	return u.articlesRepo.GetArticles(ctx, pq)
}
