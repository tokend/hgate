package cmd

import (
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/hgate/internal/conf"
	"gitlab.com/tokend/hgate/internal/router"
)

func Execute() {
	var configFile string
	var cfg conf.Config
	logger := logan.New()
	cobra.OnInitialize(func() {
		cfg = conf.NewViperConfig(configFile)
		if err := cfg.Init(); err != nil {
			logger.Error(errors.From(errors.Wrap(err, "no config provided"), logan.F{
				"config_file": configFile,
			}))
			return
		}
	})
	rootCmd := &cobra.Command{
		Use: "hgate",
		Run: func(cmd *cobra.Command, args []string) {
			r, err := router.NewRouter(cfg.Signer(), cfg.HorizonURL(), cfg.Log())
			if err != nil {
				logger.WithError(err).Error("failed to initialize router")
				return
			}
			err = http.ListenAndServe(cfg.HTTP(), r)
			if err != nil {
				logger.Error(errors.Wrap(err, "http server died"))
				return
			}
		},
	}
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml", "config file")
	rootCmd.Execute()
}
