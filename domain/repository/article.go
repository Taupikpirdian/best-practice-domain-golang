package repository

import (
	"assignee/first-domain-implement/domain/entity"
	"context"
)

type ArticleRepository interface {
	Store(ctx context.Context, dataArticle *entity.Article) error
}
