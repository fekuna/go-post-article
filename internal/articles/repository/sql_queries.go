package repository

const (
	addArticleQuery = `
		INSERT INTO articles(title, content, category, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, now(), now());
	`

	findArticleQuery = `
		SELECT * FROM articles WHERE id = ?;
	`

	updateArticleQuery = `UPDATE articles 
						SET title = COALESCE(NULLIF(?, ''), title),
						    content = COALESCE(NULLIF(?, ''), content),
						    category = COALESCE(NULLIF(?, ''), category),
						    status = COALESCE(NULLIF(?, ''), status),
						    updated_at = now()
						WHERE id = ?;
						`

	deleteUserQuery = `DELETE FROM articles WHERE id = ?`

	getTotalArticlesQuery = `SELECT COUNT(id) FROM articles;`

	getArticles = `SELECT * FROM articles ORDER BY COALESCE(NULLIF(?, ''), id) LIMIT ? OFFSET ?;`
)
