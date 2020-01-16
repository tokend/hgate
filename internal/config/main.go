package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/cop"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/tokend/connectors/keyer"
	"gitlab.com/tokend/connectors/signed"
	"gitlab.com/tokend/connectors/submit"
	"net/url"
)

var HGateRelease string

type Config interface {
	comfig.Logger
	comfig.Listenerer

	keyer.Keyer
	submit.Submitter
	signed.Clienter
	cop.Coper

	HorizonURL() *url.URL
}

type config struct {
	comfig.Logger
	comfig.Listenerer

	cop.Coper

	keyer.Keyer
	signed.Clienter
	submit.Submitter

	*horizonConfig

	getter kv.Getter
}

func NewConfig(getter kv.Getter) Config {
	return &config{
		getter:        getter,
		Logger:        comfig.NewLogger(getter, comfig.LoggerOpts{Release: HGateRelease}),
		Listenerer:    comfig.NewListenerer(getter),
		Keyer:         keyer.NewKeyer(getter),
		Submitter:     submit.NewSubmitter(getter),
		Clienter:      signed.NewClienter(getter),
		Coper:         cop.NewCoper(getter),
		horizonConfig: NewHorizon(getter),
	}
}
