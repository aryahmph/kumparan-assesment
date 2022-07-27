package pagination

import (
	"encoding/base64"
	"fmt"
	errorCommon "github.com/aryahmph/kumparan-assesment/common/error"
	"strings"
	"time"
)

func EncodeCursor(t time.Time, uuid string) string {
	key := fmt.Sprintf("%s,%s", t.Format(time.RFC3339Nano), uuid)
	return base64.StdEncoding.EncodeToString([]byte(key))
}

func DecodeCursor(encodedCursor string) (createdAt time.Time, uuid string, err error) {
	byt, err := base64.StdEncoding.DecodeString(encodedCursor)
	if err != nil {
		return createdAt, uuid, err
	}

	arrStr := strings.Split(string(byt), ",")
	if len(arrStr) != 2 {
		err = errorCommon.NewInvariantError("cursor is invalid")
		return createdAt, uuid, err
	}

	createdAt, err = time.Parse(time.RFC3339Nano, arrStr[0])
	if err != nil {
		return createdAt, uuid, err
	}
	uuid = arrStr[1]
	return createdAt, uuid, err
}
