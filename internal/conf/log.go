package conf

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/logan/v3"
)

var (
	logLevelHook = figure.Hooks{
		"map[string]string": func(value interface{}) (reflect.Value, error) {
			result, err := cast.ToStringMapStringE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse map[string]string")
			}
			return reflect.ValueOf(result), nil
		},
		"logan.Level": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				lvl, err := logan.ParseLevel(v)
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to parse log level")
				}
				return reflect.ValueOf(lvl), nil
			case nil:
				return reflect.ValueOf(nil), nil
			default:
				return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
			}
		},
	}
)

func (c *ViperConfig) Log() *logan.Entry {
	c.Lock()
	defer c.Unlock()

	if c.logan != nil {
		return c.logan
	}

	var config struct {
		Level logan.Level
	}

	err := figure.
		Out(&config).
		With(figure.BaseHooks, logLevelHook).
		From(c.GetStringMap("log")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out log"))
	}

	entry := logan.New().Level(config.Level)

	c.logan = entry
	return c.logan
}
