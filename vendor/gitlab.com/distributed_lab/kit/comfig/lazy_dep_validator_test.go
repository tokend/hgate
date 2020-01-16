package comfig

import "testing"

type withPanic struct {}
func (p withPanic) Method() string {panic("expected panic")}

type withoutPanic struct {}
func (p withoutPanic) Method() string {return ""}

type withParams struct{}
func (p withParams) Method(string) string {return ""}


func TestLazyDepValidation(t *testing.T) {
	t.Run("no panic", func(t *testing.T) {
		err := ValidateLazyDep(withoutPanic{})
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("panic", func(t *testing.T) {
		withErrors := []interface{}{withPanic{}, withParams{}}
		for _, withError := range withErrors {
			err := ValidateLazyDep(withError)
			if err == nil {
				t.Fatal("expected error not to be nil")
			}
		}

	})
}
