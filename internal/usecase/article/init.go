package article

import (
	articleRepo "github.com/aryahmph/kumparan-assesment/internal/repository/article"
	authorRepo "github.com/aryahmph/kumparan-assesment/internal/repository/author"
)

type articleUsecase struct {
	articleRepo articleRepo.Repository
	authorRepo  authorRepo.Repository
}

func NewArticleUsecase(articleRepo articleRepo.Repository, authorRepo authorRepo.Repository) Usecase {
	return articleUsecase{
		articleRepo: articleRepo,
		authorRepo:  authorRepo,
	}
}
