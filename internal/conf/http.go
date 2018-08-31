package conf

import (
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
)

func (c *ViperConfig) HTTP() string {
	c.Lock()
	defer c.Unlock()

	if c.http == nil {
		var httpConfig struct {
			Host string
			Port int
		}

		config := c.GetStringMap("http")
		if err := figure.Out(&httpConfig).From(config).Please(); err != nil {
			panic(errors.Wrap(err, "failed to figure out http"))
		}

		http := fmt.Sprintf("%s:%d", httpConfig.Host, httpConfig.Port)
		c.http = &http
	}
	return *c.http
}
