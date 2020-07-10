package figure_test

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func TestSimpleUsage(t *testing.T) {
	type Config struct {
		SomeInt                int            `fig:"some_int"`
		SomeString             string         `fig:"some_string"`
		Missing                int            `fig:"missing"`
		Default                int            `fig:"default"`
		DurationStr            time.Duration  `fig:"duration_str"`
		DurationInt            time.Duration  `fig:"duration_int"`
		DurationPointer        *time.Duration `fig:"duration_pointer"`
		DurationPointerMissing *time.Duration `fig:"duration_pointer_missing"`
		BigInt                 *big.Int       `fig:"big_int"`
		BigIntStr              *big.Int       `fig:"big_int_str"`
		Int32                  int32          `fig:"int_32"`
		Int64                  int64          `fig:"int_64"`
		Uint                   uint           `fig:"uint"`
		Uint32                 uint32         `fig:"uint_32"`
		Uint64                 uint64         `fig:"uint_64"`
		Float64                float64        `fig:"float_64"`
	}

	c := Config{
		Default: 42,
	}
	err := figure.Out(&c).From(map[string]interface{}{
		"some_int":         1,
		"some_string":      "satoshi",
		"duration_str":     "1s",
		"duration_int":     1,
		"duration_pointer": "1h",
		"big_int":          42,
		"big_int_str":      "42",
		"int_32":           16,
		"int_64":           17,
		"uint":             18,
		"uint_32":          19,
		"uint_64":          20,
		"float_64":         21.9,
	}).Please()
	if err != nil {
		t.Fatalf("expected nil error got %s", err)
	}

	duration := 1 * time.Hour
	expectedConfig := Config{
		SomeInt:         1,
		SomeString:      "satoshi",
		Default:         42,
		DurationStr:     1 * time.Second,
		DurationInt:     1,
		DurationPointer: &duration,
		BigInt:          big.NewInt(42),
		BigIntStr:       big.NewInt(42),
		Int32:           16,
		Int64:           17,
		Uint:            18,
		Uint32:          19,
		Uint64:          20,
		Float64:         21.9,
	}
	if !reflect.DeepEqual(c, expectedConfig) {
		t.Errorf("expected %#v got %#v", expectedConfig, c)
	}

}

func TestImplicitRequired(t *testing.T) {
	var config struct {
		Implicit int `fig:",required"`
	}

	err := figure.Out(&config).From(map[string]interface{}{
		"implicit": 83,
	}).Please()

	assert.NoError(t, err)
	assert.Equal(t, config.Implicit, 83)
}

func TestNonZero(t *testing.T) {
	type Config struct {
		StringVal string `fig:"string_val,non_zero"`
		IntVal int `fig:"int_val,non_zero"`
		Duration time.Duration `fig:"duration,non_zero"`
		Pointer *string `fig:"pointer,non_zero"`
	}

	pointer := "test"

	cases := []struct {
		Name     string
		Data     map[string]interface{}
		Expected Config
		Error    error
	}{
		{
			Name: "check not zero values",
			Data: map[string]interface{}{
				"string_val": "test",
				"int_val": 1,
				"duration": 1,
				"pointer": "test",
			},
			Expected: Config{
				StringVal: "test",
				IntVal: 1,
				Duration: 1,
				Pointer: &pointer,
			},
			Error: nil,
		},
		{
			Name: "check zero string values",
			Data: map[string]interface{}{
				"string_val": "",
				"int_val": 1,
				"duration": 1,
				"pointer": "test",
			},
			Expected: Config{},
			Error: figure.ErrNonZeroValue,
		},
		{
			Name: "check zero string values",
			Data: map[string]interface{}{
				"string_val": "test",
				"int_val": 0,
				"duration": 1,
				"pointer": "test",
			},
			Expected: Config{},
			Error: figure.ErrNonZeroValue,
		},
		{
			Name: "check zero duration values",
			Data: map[string]interface{}{
				"string_val": "test",
				"int_val": 1,
				"duration": 0,
				"pointer": "test",
			},
			Expected: Config{},
			Error: figure.ErrNonZeroValue,
		},
		{
			Name: "check zero pointer values",
			Data: map[string]interface{}{
				"string_val": "test",
				"int_val": 1,
				"duration": 1,
			},
			Expected: Config{},
			Error: figure.ErrNonZeroValue,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			config := Config{}
			err := figure.Out(&config).From(c.Data).Please()

			assert.Equal(t, c.Error, errors.Cause(err))
			if c.Error == nil {
				if !reflect.DeepEqual(config, c.Expected) {
					t.Errorf("expected %#v got %#v", c.Expected, config)
				}
			}
		})
	}
}

func TestNonZeroAndRequired(t *testing.T) {
	type Config struct {
		StringVal string `fig:"string_val,non_zero,required"`
	}

	cases := []struct {
		Name     string
		Data     map[string]interface{}
		Expected Config
		Error    error
	}{
		{
			Name: "check required and non_zero for same field, valid",
			Data: map[string]interface{}{
				"string_val": "test",
			},
			Expected: Config{
				StringVal: "test",
			},
			Error: nil,
		},
		{
			Name: "check required and non_zero for same field, missed in config",
			Data: map[string]interface{}{},
			Expected: Config{},
			Error: figure.ErrRequiredValue,
		},
		{
			Name: "check required and non_zero for same field, zero value in config",
			Data: map[string]interface{}{
				"string_val": "",
			},
			Expected: Config{},
			Error: figure.ErrNonZeroValue,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			config := Config{}
			err := figure.Out(&config).From(c.Data).Please()

			assert.Equal(t, c.Error, errors.Cause(err))
			if c.Error == nil {
				if !reflect.DeepEqual(config, c.Expected) {
					t.Errorf("expected %#v got %#v", c.Expected, config)
				}
			}
		})
	}
}
