package repository

import (
	"context"

	"github.com/fekuna/go-post-article/internal/articles"
	"github.com/fekuna/go-post-article/internal/models"
	"github.com/fekuna/go-post-article/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type articlesRepo struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) articles.Repository {
	return &articlesRepo{
		db: db,
	}
}

func (r *articlesRepo) AddArticle(ctx context.Context, article *models.Article) (*int64, error) {
	insertedId, err := r.db.MustExecContext(ctx, addArticleQuery, article.Title, article.Content, article.Category, article.Status).LastInsertId()
	if err != nil {
		return nil, errors.Wrap(err, "articleRepo.AddArticle.StructScan")
	}

	return &insertedId, err
}

func (r *articlesRepo) FindArticleById(ctx context.Context, articleId int) (*models.Article, error) {
	foundArticle := &models.Article{}
	if err := r.db.QueryRowxContext(ctx, findArticleQuery, articleId).StructScan(foundArticle); err != nil {
		return nil, errors.Wrap(err, "articleRepo.findArticleById.StructScan")
	}

	return foundArticle, nil
}

func (r *articlesRepo) UpdateArticle(ctx context.Context, article *models.Article) (*models.Article, error) {

	_, err := r.db.MustExecContext(ctx, updateArticleQuery, article.Title, article.Content, article.Category, article.Status, article.Id).RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "articleRepo.UpdateArticle.StructScan")
	}

	return article, err
}

func (r *articlesRepo) DeleteArticleById(ctx context.Context, articleId int) error {
	_, err := r.db.MustExecContext(ctx, deleteUserQuery, articleId).RowsAffected()
	if err != nil {
		return errors.Wrap(err, "articleRepo.UpdateArticle.StructScan")
	}

	return nil
}

func (r *articlesRepo) GetArticles(ctx context.Context, pq *utils.PaginationQuery) (*models.ArticleList, error) {
	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getTotalArticlesQuery); err != nil {
		return nil, errors.Wrap(err, "articleRepo.GetArticles.GetContext.TotalArticles")
	}

	if totalCount == 0 {
		return &models.ArticleList{
			Articles: make([]*models.Article, 0),
			Meta: models.Meta{
				TotalCount: totalCount,
				TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
				Page:       pq.GetPage(),
				Size:       pq.GetSize(),
				HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
			},
		}, nil
	}

	var articles = make([]*models.Article, 0, pq.GetSize())
	if err := r.db.SelectContext(
		ctx,
		&articles,
		getArticles,
		pq.GetOrderBy(),
		pq.GetLimit(),
		pq.GetOffset(),
	); err != nil {
		return nil, errors.Wrap(err, "articleRepo.GetArticles.SelectContext")
	}

	return &models.ArticleList{
		Articles: articles,
		Meta: models.Meta{
			TotalCount: totalCount,
			TotalPages: utils.GetTotalPages(totalCount, pq.GetSize()),
			Page:       pq.GetPage(),
			Size:       pq.GetSize(),
			HasMore:    utils.GetHasMore(pq.GetPage(), totalCount, pq.GetSize()),
		},
	}, nil
}
