package conf

import (
	"net/url"
	"sync"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/keypair"
)

type Config interface {
	Init() error
	HTTP() string
	Log() *logan.Entry
	HorizonURL() *url.URL
	Signer() keypair.Full
}

//go:generate mockery -case underscore -name rawGetter -testonly -inpkg
// rawGetter encapsulates raw config values provider
type rawGetter interface {
	GetStringMap(key string) map[string]interface{}
}

type ViperConfig struct {
	rawGetter
	*sync.RWMutex

	horizonURL *url.URL
	logan      *logan.Entry
	http       *string
	signer     keypair.Full
}

func NewViperConfig(fn string) Config {
	// init underlying viper
	v := viper.GetViper()
	v.SetConfigFile(fn)

	return newViperConfig(v)
}

func newViperConfig(raw rawGetter) Config {
	config := &ViperConfig{
		RWMutex: &sync.RWMutex{},
	}
	config.rawGetter = raw
	return config
}

func (c *ViperConfig) Init() error {
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	return nil
}
