package comfig

import (
	"io"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnce_Do(t *testing.T) {
	t.Run("run once", func(t *testing.T) {
		var once Once
		var wg sync.WaitGroup
		runs := make(chan struct{}, 2<<10)
		for i := 0; i < 2<<10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				once.Do(func() interface{} {
					runs <- struct{}{}
					return 1
				})

			}()
		}
		wg.Wait()
		assert.Len(t, runs, 1)
	})

	t.Run("value passed", func(t *testing.T) {
		var once Once
		expected := 1
		got := once.Do(func() interface{} {
			return expected
		})
		assert.Equal(t, expected, got)
	})

	t.Run("value persists", func(t *testing.T) {
		var once Once
		expected := 1
		once.Do(func() interface{} { return expected })
		got := once.Do(func() interface{} { panic("should not happen") })
		assert.Equal(t, expected, got)
	})

	t.Run("panic passed", func(t *testing.T) {
		var once Once
		assert.PanicsWithValue(t, io.EOF, func() {
			once.Do(func() interface{} {
				panic(io.EOF)
			})
		})
	})

	t.Run("panic persists", func(t *testing.T) {
		var once Once
		assert.PanicsWithValue(t, io.EOF, func() {
			once.Do(func() interface{} {
				panic(io.EOF)
			})
		})
		assert.PanicsWithValue(t, io.EOF, func() {
			once.Do(func() interface{} {
				return 1
			})
		})
	})
}
