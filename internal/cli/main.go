package cli

import (
	"context"
	"github.com/tokend/hgate/internal/config"
	"github.com/tokend/hgate/internal/service"
	"github.com/urfave/cli"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func Run(args []string) bool {
	var cfg config.Config
	log := logan.New()

	defer func() {
		if rvr := recover(); rvr != nil {
			log.WithRecover(rvr).Error("app panicked")
		}
	}()

	app := cli.NewApp()

	before := func(_ *cli.Context) error {
		getter, err := kv.FromEnv()
		if err != nil {
			return errors.Wrap(err, "failed to get config")
		}
		cfg = config.NewConfig(getter)
		log = cfg.Log()
		return nil
	}

	app.Commands = cli.Commands{
		{
			Name:   "run",
			Before: before,
			Action: func(_ *cli.Context) error {
				s := service.New(cfg)
				return s.Run(context.Background(), cfg)
			},
		},
	}

	if err := app.Run(args); err != nil {
		log.WithError(err).Error("app finished")
		return false
	}
	return true
}
