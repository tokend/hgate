package cop

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func TestChiRules(t *testing.T) {
	c := Cop{}

	t.Run("empty router", func(t *testing.T) {
		r := chi.NewRouter()

		rule, err := c.GetRule(r)
		assert.NoError(t, err)
		assert.Len(t, rule, 0)
	})

	t.Run("flat", func(t *testing.T) {
		r := chi.NewRouter()
		r.Get("/a", nil)
		r.Post("/a/{id}/b", nil)

		rule, err := c.GetRule(r)
		assert.NoError(t, err)

		assert.Equal(t, "(Path(`/a`)&&Method(`GET`))||(Path(`/a/{id}/b`)&&Method(`POST`))", rule)
	})

	t.Run("nested", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/a", func(r chi.Router) {
			r.Get("/", nil)
			r.Post("/{id}/b", nil)
		})

		rule, err := c.GetRule(r)
		assert.NoError(t, err)

		assert.Equal(t, "(Path(`/a`)&&Method(`GET`))||(Path(`/a/{id}/b`)&&Method(`POST`))", rule)
	})

	cWithPrefix := New(CopConfig{
		ServicePrefix: "/test-prefix",
	})
	t.Run("with prefix", func(t *testing.T) {
		r := chi.NewRouter()
		r.Route("/test-prefix", func(r chi.Router) {
			r.Get("/", nil)
			r.Post("/{id}/b", nil)
		})

		rule, err := cWithPrefix.GetRule(r)
		assert.NoError(t, err)

		assert.Equal(t, "PathPrefix(`/test-prefix`)", rule)
	})
}
