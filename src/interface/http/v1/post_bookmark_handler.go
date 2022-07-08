package v1

import (
	"database/sql"
	"github.com/edwintantawi/gomark/src/common/helper"
	"github.com/edwintantawi/gomark/src/domain/bookmark"
	"github.com/edwintantawi/gomark/src/domain/response"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type postBookmarkHandler struct {
	useCase bookmark.AddBookmarkUseCase
}

func NewPostBookmarkHandler(useCase bookmark.AddBookmarkUseCase) *postBookmarkHandler {
	return &postBookmarkHandler{useCase: useCase}
}

func (h *postBookmarkHandler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var newBookmark bookmark.New
	helper.BodyParser(r.Body, &newBookmark)

	txHandle := r.Context().Value("dbTx").(*sql.Tx)

	newBookmark.Validate()
	bookmarkId := h.useCase.WithTx(txHandle).Execute(r.Context(), newBookmark)

	res := response.Body{
		Status:  http.StatusText(http.StatusCreated),
		Message: "Bookmark added successfully",
		Result:  bookmarkId,
	}

	helper.ResponseJson(w, http.StatusCreated, res)
}
