package figure_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func TestWithoutHook(t *testing.T) {
	type SomeStruct struct {
		AnotherInt    int    `fig:"another_int"`
		AnotherString string `fig:"another_string"`
	}

	type Config struct {
		SomeInt    int        `fig:"some_int"`
		SomeString string     `fig:"some_string"`
		SomeStruct SomeStruct `fig:"some_struct"`
		MoreInt    int        `fig:"more_int"`
	}

	c := Config{}

	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":    1,
		"some_string": "satoshi",
		"some_struct": map[string]interface{}{"another_int": 5, "another_string": "another"},
		"more_int":    7,
	}).Please()
	if err != nil {
		t.Fatalf("expected nil error got %s", err)
	}

	expectedConfig := Config{
		SomeInt:    1,
		SomeString: "satoshi",
		SomeStruct: SomeStruct{AnotherInt: 5, AnotherString: "another"},
		MoreInt:    7,
	}
	assert.EqualValues(t, expectedConfig, c)
}

func TestPointerStruct(t *testing.T) {
	type SomeStruct struct {
		AnotherInt    int    `fig:"another_int"`
		AnotherString string `fig:"another_string"`
	}

	type Config struct {
		SomeInt     int         `fig:"some_int"`
		SomeString  string      `fig:"some_string"`
		SomeStructP *SomeStruct `fig:"some_struct"`
		MoreInt     int         `fig:"more_int"`
	}

	c := Config{}

	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":    1,
		"some_string": "satoshi",
		"some_struct": map[string]interface{}{"another_int": 5, "another_string": "another"},
		"more_int":    7,
	}).Please()

	assert.Error(t, err)
	assert.Equal(t, figure.ErrNoHook, errors.Cause(err))

}
func TestWithRequired(t *testing.T) {
	type AnotherStruct struct {
		anInt    int    `fig:"an_int,required"`
		anString string `fig:"an_string"`
	}

	type SomeStruct struct {
		AnotherInt    int           `fig:"another_int"`
		AnotherString string        `fig:"another_string"`
		AnotherStruct AnotherStruct `fig:"another_struct,required"`
	}

	type Config struct {
		SomeInt    int        `fig:"some_int"`
		SomeString string     `fig:"some_string"`
		SomeStruct SomeStruct `fig:"some_struct"`
		MoreInt    int        `fig:"more_int"`
	}

	c := Config{}

	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":    1,
		"some_string": "satoshi",
		"some_struct": map[string]interface{}{"another_int": 5, "another_string": "another"},
		"more_int":    7,
	}).Please()

	assert.Error(t, err)
	assert.Equal(t, figure.ErrRequiredValue, errors.Cause(err))
}

func TestWithInternalRequired(t *testing.T) {
	type AnotherStruct struct {
		anInt    int    `fig:"an_int,required"`
		anString string `fig:"an_string"`
	}

	type SomeStruct struct {
		AnotherInt    int           `fig:"another_int"`
		AnotherString string        `fig:"another_string"`
		AnotherStruct AnotherStruct `fig:"another_struct"`
	}

	type Config struct {
		SomeInt    int        `fig:"some_int"`
		SomeString string     `fig:"some_string"`
		SomeStruct SomeStruct `fig:"some_struct"`
		MoreInt    int        `fig:"more_int"`
	}

	c := Config{}

	InternalStruct := AnotherStruct{anString: "only one"}

	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":    1,
		"some_string": "satoshi",
		"some_struct": map[string]interface{}{"another_int": 5, "another_string": "another", "another_struct": InternalStruct},
		"more_int":    7,
	}).Please()

	assert.Error(t, err)
	assert.Equal(t, figure.ErrRequiredValue, errors.Cause(err))
}

func TestManyStructs(t *testing.T) {
	type ThirdStruct struct {
		ThirdInt    int    `fig:"third_int"`
		ThirdString string `fig:"third_string"`
	}
	type SecondStruct struct {
		SecondtInt   int         `fig:"second_int"`
		SecondString string      `fig:"second_string"`
		Struct       ThirdStruct `fig:"third_struct"`
	}
	type FirstStruct struct {
		FirstInt    int          `fig:"first_int"`
		FirstString string       `fig:"first_string"`
		Struct      SecondStruct `fig:"second_struct"`
	}

	type Config struct {
		SomeInt    int         `fig:"some_int"`
		SomeString string      `fig:"some_string"`
		SomeStruct FirstStruct `fig:"first_struct"`
		MoreInt    int         `fig:"more_int"`
	}

	c := Config{}

	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":    1,
		"some_string": "one",
		"first_struct": map[string]interface{}{"first_int": 5, "first_string": "first",
			"second_struct": map[string]interface{}{"second_int": 2, "second_string": "second",
				"third_struct": map[string]interface{}{"third_int": 3, "third_string": "third"}}},
		"more_int": 7,
	}).Please()
	if err != nil {
		t.Fatalf("expected nil error got %s", err)
	}

	expectedConfig := Config{
		SomeInt:    1,
		SomeString: "one",
		SomeStruct: FirstStruct{FirstInt: 5, FirstString: "first",
			Struct: SecondStruct{SecondtInt: 2, SecondString: "second",
				Struct: ThirdStruct{ThirdInt: 3, ThirdString: "third"}}},
		MoreInt: 7,
	}
	assert.EqualValues(t, expectedConfig, c)
}

func TestNotValidStruct(t *testing.T) {
	type AnotherStruct struct {
		anInt    int    `fig:"an_int,required"`
		anString string `fig:"an_string"`
	}

	type SomeStruct struct {
		AnotherInt    int           `fig:"another_int"`
		AnotherString string        `fig:"another_string"`
		AnotherStruct AnotherStruct `fig:"another_struct,required"`
	}

	type Config struct {
		SomeInt    int        `fig:"some_int"`
		SomeString string     `fig:"some_string"`
		SomeStruct SomeStruct `fig:"some_struct"`
		MoreInt    int        `fig:"more_int"`
	}

	c := Config{}

	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":    1,
		"some_string": "satoshi",
		"some_struct": map[string]interface{}{"another_int": 5, "another_string": "another"},
		"more_int":    7,
	}).Please()

	assert.Error(t, err)
	assert.Equal(t, figure.ErrRequiredValue, errors.Cause(err))
}

func TestStructFigure(t *testing.T) {
	type SomeStruct struct {
		AnotherInt int `fig:"another_int"`
	}

	type Config struct {
		SomeStruct SomeStruct `fig:"some_struct"`
	}
	t.Run("invalid value", func(t *testing.T) {
		c := Config{}
		err := figure.Out(&c).From(map[string]interface{}{
			"some_struct": 7,
		}).Please()

		assert.Error(t, err)
		assert.Equal(t, figure.ErrNotValid, errors.Cause(err))

	})
}
