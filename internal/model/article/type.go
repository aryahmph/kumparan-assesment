package article

import (
	"github.com/aryahmph/kumparan-assesment/internal/model"
	authorModel "github.com/aryahmph/kumparan-assesment/internal/model/author"
)

type (
	Article struct {
		ID     string
		Title  string
		Body   string
		Author Author

		Timestamp
	}

	Author    = authorModel.Author
	Timestamp = model.Timestamp
)
