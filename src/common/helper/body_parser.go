package helper

import (
	"encoding/json"
	"github.com/edwintantawi/gomark/src/common/exception"
	"io"
)

func BodyParser(body io.ReadCloser, target any) {
	r, _ := io.ReadAll(body)

	if err := json.Unmarshal(r, &target); err != nil {
		PanicError(exception.NewInvariantError("Invalid request body"))
	}
}
