package internal

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi"
)

func TestChi(t *testing.T) {
	t.Run("empty router", func(t *testing.T) {
		r := chi.NewRouter()

		services, err := NewChi(r).Services()
		assert.NoError(t, err)
		assert.Len(t, services, 0)
	})

	t.Run("flat", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/a", nil)
		r.Post("/a/{id}/b", nil)

		services, err := NewChi(r).Services()
		assert.NoError(t, err)
		assert.Len(t, services, 2)

		assert.Equal(t, "get-a", services[0].Name)
		assert.Len(t, services[0].Proxy.Methods, 1)
		assert.Equal(t, http.MethodGet, services[0].Proxy.Methods[0])
		assert.Equal(t, "/a", services[0].Proxy.ListenPath)

		assert.Equal(t, "post-a-x-b", services[1].Name)
		assert.Len(t, services[1].Proxy.Methods, 1)
		assert.Equal(t, http.MethodPost, services[1].Proxy.Methods[0])
		assert.Equal(t, "/a/{id}/b", services[1].Proxy.ListenPath)
	})

	t.Run("nested", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/a", func(r chi.Router) {
			r.Get("/", nil)
			r.Post("/{id}/b", nil)
		})

		services, err := NewChi(r).Services()
		assert.NoError(t, err)
		assert.Len(t, services, 2)

		assert.Equal(t, "get-a", services[0].Name)
		assert.Len(t, services[0].Proxy.Methods, 1)
		assert.Equal(t, http.MethodGet, services[0].Proxy.Methods[0])
		assert.Equal(t, "/a", services[0].Proxy.ListenPath)

		assert.Equal(t, "post-a-x-b", services[1].Name)
		assert.Len(t, services[1].Proxy.Methods, 1)
		assert.Equal(t, http.MethodPost, services[1].Proxy.Methods[0])
		assert.Equal(t, "/a/{id}/b", services[1].Proxy.ListenPath)
	})
}
