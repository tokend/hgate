package keyer

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair"
	"gitlab.com/tokend/keypair/figurekeypair"
)

type Keys struct {
	Signer keypair.Full    `fig:"signer"`
	Source keypair.Address `fig:"source"`
}

type Keyer interface {
	Keys() Keys
}

type keyer struct {
	once   comfig.Once
	getter kv.Getter
}

func NewKeyer(getter kv.Getter) *keyer {
	return &keyer{
		getter: getter,
	}
}

func (k *keyer) Keys() Keys {
	return k.once.Do(func() interface{} {
		var config Keys
		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(k.getter, "keys")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out api"))
		}

		return config
	}).(Keys)
}