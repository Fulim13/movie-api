package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValues := fmt.Sprintf("%d mins", r)
	quotedJsonValues := strconv.Quote(jsonValues)
	return []byte(quotedJsonValues), nil
}
