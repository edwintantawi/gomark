package bookmark

import (
	"github.com/edwintantawi/gomark/src/common/exception"
	"time"
)

type ID string

type Added struct {
	ID          ID        `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type New struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

func (b *New) Validate() {
	if b.Title == "" {
		panic(exception.NewInvariantError("Title is required"))
	}

	if b.Description == "" {
		panic(exception.NewInvariantError("Description is required"))
	}

	if b.Url == "" {
		panic(exception.NewInvariantError("Url is required"))
	}
}
