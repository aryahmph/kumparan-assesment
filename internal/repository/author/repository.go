package author

import (
	"context"
	authorModel "github.com/aryahmph/kumparan-assesment/internal/model/author"
)

type Repository interface {
	FindByID(ctx context.Context, id string) (author authorModel.Author, err error)
}
