package http

import (
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	articleModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

func (d HTTPArticleDelivery) mapArticleModelToResponse(model articleModel.Article) httpCommon.Article {
	return httpCommon.Article{
		ID:    model.ID,
		Title: model.Title,
		Body:  model.Body,
		Author: httpCommon.Author{
			ID:   model.Author.ID,
			Name: model.Author.Name,
		},
		CreatedAt: model.CreatedAt,
	}
}

func (d HTTPArticleDelivery) mapArticleBodyToModel(payload httpCommon.AddArticle) articleModel.Article {
	return articleModel.Article{
		Title:  payload.Title,
		Body:   payload.Body,
		Author: articleModel.Author{ID: payload.AuthorID},
	}
}
