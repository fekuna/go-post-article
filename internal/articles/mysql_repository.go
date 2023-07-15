package articles

import (
	"context"

	"github.com/fekuna/go-post-article/internal/models"
	"github.com/fekuna/go-post-article/pkg/utils"
)

type Repository interface {
	AddArticle(ctx context.Context, article *models.Article) (*int64, error)
	FindArticleById(ctx context.Context, articleId int) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	DeleteArticleById(ctx context.Context, articleId int) error
	GetArticles(ctx context.Context, pq *utils.PaginationQuery) (*models.ArticleList, error)
}
