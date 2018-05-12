package config

import (
	"errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/swarmfund/go/keypair"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
)

type GateConfig struct {
	Port       string `yaml:"port"`
	HorizonUrl string `yaml:"horizon_url"`
	Seed       string `yaml:"seed"`
	LogLevel   string `yaml:"log_level"`

	HUrl *url.URL
	KP   *keypair.Full
	LL   logan.Level
}

func InitConfig(filePath string) (*GateConfig, error) {
	rawConfig, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config = new(GateConfig)
	err = yaml.Unmarshal(rawConfig, config)
	if err != nil {
		return nil, err
	}

	config.HUrl, err = url.Parse(config.HorizonUrl)
	if err != nil {
		return nil, err
	}

	err = config.parseKP()
	if err != nil {
		return nil, err
	}

	config.LL = logLevel(config.LogLevel)
	return config, nil
}

func (gc *GateConfig) parseKP() error {
	kp, err := keypair.Parse(gc.Seed)
	if err != nil {
		return err
	}

	var ok bool
	gc.KP, ok = kp.(*keypair.Full)
	if !ok {
		return errors.New("must be a seed")
	}
	return nil
}

func logLevel(ll string) logan.Level {
	switch ll {
	case "debug":
		return logan.DebugLevel
	case "error":
		return logan.ErrorLevel
	case "info":
		return logan.InfoLevel
	case "warn":
		return logan.WarnLevel
	default:
		return logan.WarnLevel
	}
}
