package article

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	httpCommon "github.com/aryahmph/kumparan-assesment/common/http"
	paginationCommon "github.com/aryahmph/kumparan-assesment/common/pagination"
	articleModel "github.com/aryahmph/kumparan-assesment/internal/model/article"
)

func (usecase articleUsecase) GetAll(ctx context.Context, pParams httpCommon.PaginationQueryParams, aParams httpCommon.ArticleQueryParams) (articles []articleModel.Article, err error) {
	// Generate SQL statement from query params filter
	queryBuilder := sq.Select("articles.id", "title", "body", "articles.created_at", "authors.id", "authors.name").
		From("articles").Join("authors ON articles.author_id = authors.id").
		OrderBy("articles.created_at DESC, articles.id DESC").Limit(uint64(pParams.Limit)).
		PlaceholderFormat(sq.Dollar)

	if aParams != (httpCommon.ArticleQueryParams{}) {
		// WHERE text_search @@ plainto_tsquery('indonesian', $x) ORDER BY ts_rank(text_search, plainto_tsquery('indonesian', $x)) DESC
		if aParams.Keyword != "" {
			queryBuilder = queryBuilder.Where("text_search @@ plainto_tsquery('indonesian', ?)", aParams.Keyword).
				OrderByClause("ts_rank(text_search, plainto_tsquery('indonesian', ?)) DESC", aParams.Keyword)
		}

		// WHERE name_search @@ plainto_tsquery('indonesian', $x) ORDER BY ts_rank(name_search, plainto_tsquery('indonesian', $x)) DESC
		if aParams.Author != "" {
			queryBuilder = queryBuilder.Where("name_search @@ plainto_tsquery('indonesian', ?)", aParams.Author).
				OrderByClause("ts_rank(name_search, plainto_tsquery('indonesian', ?)) DESC", aParams.Author)
		}
	}

	if pParams != (httpCommon.PaginationQueryParams{}) {
		if pParams.Cursor != "" {
			createdAt, uuid, err := paginationCommon.DecodeCursor(pParams.Cursor)
			if err != nil {
				return articles, err
			}

			// WHERE articles.created_at <= $x AND articles.id < $y
			queryBuilder = queryBuilder.Where(sq.And{
				sq.LtOrEq{"articles.created_at": createdAt},
				sq.Lt{"articles.id": uuid},
			})
		}
	}

	statement, args, err := queryBuilder.ToSql()
	if err != nil {
		return articles, err
	}

	return usecase.articleRepo.FindAll(ctx, statement, args)
}
