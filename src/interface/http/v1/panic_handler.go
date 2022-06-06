package v1

import (
	"github.com/edwintantawi/gomark/src/common/exception"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/domain/response"
	"net/http"
)

type panicHandler struct{}

func NewPanicHandler() *panicHandler {
	return &panicHandler{}
}

func (h *panicHandler) Handle(w http.ResponseWriter, _ *http.Request, errorx interface{}) {
	if err, ok := errorx.(exception.Error); ok {
		res := response.Body{
			Message: err.Error(),
			Status:  http.StatusText(err.Code()),
			Result:  nil,
		}
		helper.ResponseJson(w, err.Code(), res)
		return
	}

	err := errorx.(error)
	res := response.Body{
		Message: err.Error(),
		Status:  http.StatusText(http.StatusInternalServerError),
		Result:  nil,
	}
	helper.ResponseJson(w, http.StatusInternalServerError, res)

}
