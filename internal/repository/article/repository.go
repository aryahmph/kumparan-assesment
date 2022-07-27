package article

import (
	"context"
	articleModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

type Repository interface {
	Insert(ctx context.Context, article articleModel.Article) (id string, err error)
	FindAll(ctx context.Context, statement string, args []interface{}) (articles []articleModel.Article, err error)
}
