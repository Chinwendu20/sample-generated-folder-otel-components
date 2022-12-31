
package project

import (
	"context"
	"go.opentelemetry.io/collector/extension"
)

const (
	// typeStr is the type of the extension
	typeStr = "project"
)


// NewFactory creates a  extension factory
func NewFactory() extension.Factory {
	return extension.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the extension type that you would like, change the second parameter as you like
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		// extension.WithMetrics(f.createMetricsExtension, component.StabilityLevelBeta),
		// extension.WithTraces(f.createTracesExtension, component.StabilityLevelBeta),
		// extension.WithLogs(f.createLogsExtension, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &Config{

	}
}

// createMetricsExtension creates a metrics extension based on this config.
func createMetricsExtension(
	ctx context.Context,
	set extension.CreateSettings,
	c component.Config,
) (extension.Metrics, error) {

	
}

// createTracesExtension creates a trace extension based on this config.
func  createTracesExtension(
	ctx context.Context,
	set extension.CreateSettings,
	c component.Config,
) (extension.Traces, error) {

}

func  createLogsExtension(
	ctx context.Context,
	set extension.CreateSettings,
	c component.Config,
) (extension.Logs, error) {

}