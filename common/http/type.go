package http

import "time"

type (
	Response struct {
		Metadata `json:"_metadata,omitempty"`
		Data     any `json:"data"`
	}

	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	Author struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Article struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		Author    Author    `json:"_author"`
	}

	AddArticle struct {
		Title    string `json:"title" binding:"required,lte=255"`
		Body     string `json:"body" binding:"required"`
		AuthorID string `json:"author_id" binding:"required,uuid4"`
	}

	ArticleQueryParams struct {
		Keyword string `form:"query" json:"keyword"`
		Author  string `form:"author" json:"author"`
	}

	Metadata struct {
		Limit  int    `json:"limit"`
		Count  int    `json:"count"`
		Cursor Cursor `json:"cursor"`
	}

	Cursor struct {
		Self string `json:"self"`
		Next string `json:"next"`
	}

	PaginationQueryParams struct {
		Limit  int    `form:"limit" json:"limit" binding:"gte=0"`
		Cursor string `form:"cursor" json:"cursor" binding:"omitempty,base64"`
	}
)
