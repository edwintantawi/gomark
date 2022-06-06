package helper

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockTX struct {
	mock.Mock
}

func (m *mockTX) Rollback() error {
	args := m.Called()
	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func (m *mockTX) Commit() error {
	args := m.Called()
	if args[0] != nil {
		return args[0].(error)
	}

	return nil
}

func TestHandleDeferTX(t *testing.T) {
	t.Run("it should tx rollback when panic, and not panic when rollback not error", func(t *testing.T) {
		tx := new(mockTX)
		tx.On("Rollback").Return(nil)
		tx.On("Commit").Return(nil)

		defer func() {
			// expect rollback to not throw panic
			err := recover()
			assert.Nil(t, err)
		}()

		defer func() {
			tx.AssertNumberOfCalls(t, "Rollback", 1)
			tx.AssertNumberOfCalls(t, "Commit", 0)
		}()

		defer HandleDeferTX(tx)

		panic("an panic attack")
	})

	t.Run("it should tx rollback when panic, and panic when rollback error", func(t *testing.T) {
		tx := new(mockTX)
		tx.On("Rollback").Return(errors.New("rollback error"))
		tx.On("Commit").Return(nil)

		defer func() {
			// expect rollback to throw panic
			err := recover()
			assert.NotNil(t, err)
		}()

		defer func() {
			tx.AssertNumberOfCalls(t, "Rollback", 1)
			tx.AssertNumberOfCalls(t, "Commit", 0)
		}()

		defer HandleDeferTX(tx)

		panic("an panic attack")
	})

	t.Run("it should tx commit when not panic, and not panic when commit not error", func(t *testing.T) {
		tx := new(mockTX)
		tx.On("Rollback").Return(nil)
		tx.On("Commit").Return(nil)

		defer func() {
			// expect commit to not throw panic
			err := recover()
			assert.Nil(t, err)
		}()

		defer func() {
			tx.AssertNumberOfCalls(t, "Rollback", 0)
			tx.AssertNumberOfCalls(t, "Commit", 1)
		}()

		defer HandleDeferTX(tx)
	})

	t.Run("it should tx commit when not panic, and panic when commit error", func(t *testing.T) {
		tx := new(mockTX)
		tx.On("Rollback").Return(nil)
		tx.On("Commit").Return(errors.New("commit error"))

		defer func() {
			// expect commit to not throw panic
			err := recover()
			assert.NotNil(t, err)
		}()

		defer func() {
			tx.AssertNumberOfCalls(t, "Rollback", 0)
			tx.AssertNumberOfCalls(t, "Commit", 1)
		}()

		defer HandleDeferTX(tx)
	})
}
