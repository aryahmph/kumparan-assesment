package article

import (
	"context"
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	articleModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

type Usecase interface {
	GetAll(ctx context.Context, pParams httpCommon.PaginationQueryParams, aParams httpCommon.ArticleQueryParams) (articles []articleModel.Article, err error)
	Create(ctx context.Context, article articleModel.Article) (id string, err error)
}
