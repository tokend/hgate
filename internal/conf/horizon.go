package conf

import (
	"net/url"
	"reflect"

	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gitlab.com/distributed_lab/figure"
)

var (
	// TODO: move to figure
	URLHook = figure.Hooks{
		"url.URL": func(value interface{}) (reflect.Value, error) {
			str, err := cast.ToStringE(value)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse string")
			}
			u, err := url.Parse(str)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to parse url")
			}
			return reflect.ValueOf(*u), nil
		},
	}
)

func (c *ViperConfig) HorizonURL() *url.URL {
	c.Lock()
	defer c.Unlock()

	if c.horizonURL == nil {
		var config struct {
			URL url.URL
		}

		if err := figure.Out(&config).From(c.GetStringMap("horizon")).With(URLHook).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out horizon"))
		}

		c.horizonURL = &config.URL
	}
	return c.horizonURL
}
