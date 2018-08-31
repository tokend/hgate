package conf

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/tokend/keypair"
)

var (
	// TODO: move to figure
	KeypairHook = figure.Hooks{
		"keypair.Full": func(value interface{}) (reflect.Value, error) {
			switch v := value.(type) {
			case string:
				kp, err := keypair.ParseSeed(v)
				if err != nil {
					return reflect.Value{}, errors.Wrap(err, "failed to parse kp")
				}
				return reflect.ValueOf(kp), nil
			case nil:
				return reflect.ValueOf(nil), nil
			default:
				return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
			}
		},
	}
)

func (c *ViperConfig) Signer() keypair.Full {
	c.Lock()
	defer c.Unlock()

	if c.signer == nil {
		var config struct {
			Signer keypair.Full
		}

		if err := figure.Out(&config).From(c.GetStringMap("keypair")).With(KeypairHook).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out keypair"))
		}

		c.signer = config.Signer
	}
	return c.signer
}
