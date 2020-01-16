package service

import (
	"context"
	"github.com/tokend/hgate/internal/config"
	"github.com/tokend/hgate/internal/web"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/connectors/lazyinfo"
	"gitlab.com/tokend/go/xdrbuild"
	"net/http"
)

type Service struct {
	log *logan.Entry
}

func New(cfg config.Config) *Service {
	return &Service{
		log: cfg.Log(),
	}
}

func (s *Service) Run(ctx context.Context, cfg config.Config) error {
	infoer := lazyinfo.New(cfg.Client())

	info, err := infoer.Info()
	if err != nil {
		return errors.Wrap(err, "failed to get horizon info")
	}

	builder := xdrbuild.NewBuilder(info.Attributes.NetworkPassphrase, info.Attributes.TxExpirationPeriod)

	r := web.Router(cfg, builder)

	err = cfg.Cop().RegisterChi(r)
	if err != nil {
		return errors.Wrap(err, "failed to register service")
	}

	err = http.Serve(cfg.Listener(), r)
	if err != nil {
		return errors.Wrap(err, "server stopped with error")
	}

	return nil
}
