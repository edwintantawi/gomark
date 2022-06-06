package bookmark

import (
	"github.com/edwintantawi/gomark/src/common/exception"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBookmarkValidate(t *testing.T) {
	t.Run("it should panic invariant exception if Title is empty", func(t *testing.T) {
		newBookmark := New{
			Description: "an description",
			Url:         "https://example.com",
		}

		e := exception.NewInvariantError("Title is required")

		assert.PanicsWithValue(t, e, newBookmark.Validate)
	})

	t.Run("it should panic invariant exception if Description is empty", func(t *testing.T) {
		newBookmark := New{
			Title: "GoMark",
			Url:   "https://example.com",
		}

		e := exception.NewInvariantError("Description is required")

		assert.PanicsWithValue(t, e, newBookmark.Validate)
	})

	t.Run("it should panic invariant exception if Url is empty", func(t *testing.T) {
		newBookmark := New{
			Title:       "GoMark",
			Description: "an description",
		}

		e := exception.NewInvariantError("Url is required")

		assert.PanicsWithValue(t, e, newBookmark.Validate)
	})

	t.Run("it should not panic invariant exception if all data exist", func(t *testing.T) {
		newBookmark := New{
			Title:       "GoMark",
			Description: "an description",
			Url:         "https://example.com",
		}

		assert.NotPanics(t, newBookmark.Validate)
	})
}
