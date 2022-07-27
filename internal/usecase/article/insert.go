package article

import (
	"context"
	articleModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

func (usecase articleUsecase) Create(ctx context.Context, article articleModel.Article) (id string, err error) {
	_, err = usecase.authorRepo.FindByID(ctx, article.Author.ID)
	if err != nil {
		return id, err
	}
	return usecase.articleRepo.Insert(ctx, article)
}
