package author

import "github.com/aryahmph/kumparan-assesment/internal/model"

type (
	Author struct {
		ID   string
		Name string

		Timestamp
	}

	Timestamp = model.Timestamp
)
