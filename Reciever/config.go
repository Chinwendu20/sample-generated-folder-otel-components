
package project


import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
)


// Config defines configuration for your receiver.
type Config struct {
	
}

var _ component.Config = (*Config)(nil)
var _ confmap.Unmarshaler = (*Config)(nil)

// Validate the configuration for errors to implement the configvalidator interface.
// You can skip this if you do not want to validate your config
func (c *Config) Validate() error {
	return nil

}



// Specify how to unmarshal your config, you can skip this if you want to make use of the default unmarshaller.
func (c *Config) Unmarshal(configMap *confmap.Conf) error {
	return nil
}
