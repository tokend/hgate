package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetName(t *testing.T) {
	t.Run("with parameter", func(t *testing.T) {
		endpoint := "/users/{id}"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-users-x", name)
	})

	t.Run("simple", func(t *testing.T) {
		endpoint := "/users"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-users", name)
	})

	t.Run("root", func(t *testing.T) {
		endpoint := "/"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get", name)
	})

	t.Run("with two parameters", func(t *testing.T) {
		endpoint := "/users/{id}/another/{one}"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-users-x-another-x", name)
	})

	t.Run("without parameters", func(t *testing.T) {
		endpoint := "/another/one/and/another/one"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-another-one-and-another-one", name)
	})

	t.Run("with underscore", func(t *testing.T) {
		endpoint := "/another_one"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-another-one", name)
	})

	t.Run("parameter with underscore", func(t *testing.T) {
		endpoint := "/users/{user_id}/another/{one}"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-users-x-another-x", name)
	})

	t.Run("few parameters", func(t *testing.T) {
		endpoint := "/order_books/{base}:{quote}:{order_book_id}/something"
		method := "GET"
		name := GetName(endpoint, method)
		assert.Equal(t, "get-order-books-x-x-x-something", name)
	})
}
