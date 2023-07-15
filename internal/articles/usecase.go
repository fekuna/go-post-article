package articles

import (
	"context"

	"github.com/fekuna/go-post-article/internal/models"
	"github.com/fekuna/go-post-article/pkg/utils"
)

type UseCase interface {
	AddArticle(ctx context.Context, user *models.Article) (*int64, error)
	FindArticleById(ctx context.Context, articleId int) (*models.Article, error)
	UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error)
	DeleteArticle(ctx context.Context, articleId int) error
	GetArticles(ctx context.Context, pq *utils.PaginationQuery) (*models.ArticleList, error)
}
