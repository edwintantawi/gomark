package v1

import (
	"database/sql"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/edwintantawi/gomark/src/domain/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getBookmarksHandler struct {
	useCase bookmark.GetBookmarksUseCase
}

func NewGetBookmarksHandler(useCase bookmark.GetBookmarksUseCase) *getBookmarksHandler {
	return &getBookmarksHandler{useCase: useCase}
}

func (h *getBookmarksHandler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	txHandle := r.Context().Value("dbTx").(*sql.Tx)

	bookmarks := h.useCase.WithTx(txHandle).Execute(r.Context())

	res := response.Body{
		Status:  http.StatusText(http.StatusOK),
		Message: "All Bookmarks",
		Result:  bookmarks,
	}

	helper.ResponseJson(w, http.StatusOK, res)
}
