package signed

import (
	"gitlab.com/tokend/connectors/keyer"
	"net/http"
	"net/url"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/keypair/figurekeypair"
)

type Clienter interface {
	Client() *Client
}

type clienter struct {
	getter kv.Getter
	keyer.Keyer
	once comfig.Once
}

func NewClienter(getter kv.Getter) *clienter {
	return &clienter{
		getter: getter,
		Keyer:  keyer.NewKeyer(getter),
	}
}

func (h *clienter) Client() *Client {
	return h.once.Do(func() interface{} {
		keys := h.Keyer.Keys()

		var config struct {
			Endpoint *url.URL `fig:"endpoint,required"`
		}

		err := figure.
			Out(&config).
			With(figure.BaseHooks, figurekeypair.Hooks).
			From(kv.MustGetStringMap(h.getter, "client")).
			Please()
		if err != nil {
			panic(errors.Wrap(err, "failed to figure out client"))
		}

		cli := NewClient(http.DefaultClient, config.Endpoint)
		if keys.Signer != nil {
			cli = cli.WithSigner(keys.Signer)
		}

		return cli
	}).(*Client)
}
