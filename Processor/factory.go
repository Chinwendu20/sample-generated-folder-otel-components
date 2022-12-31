
package project

import (
	"context"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/component"


)

const (
	// typeStr is the type of the processor
	typeStr = "project"
)


// NewFactory creates a processor factory
func NewFactory() processor.Factory {
	return processor.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the processor type that you would like, change the second parameter as you like
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		// processor.WithMetrics(createMetricsProcessor, component.StabilityLevelBeta),
		// processor.WithTraces(createTracesProcessor, component.StabilityLevelBeta),
		// processor.WithLogs(createLogsProcessor, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &Config{

	}
}

// createMetricsProcessor creates a metrics processor based on this config.
func createMetricsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	c component.Config,
) (processor.Metrics, error) {

	
}

// createTracesProcesor creates a trace processor based on this config.
func  createTracesProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	c component.Config,
) (processor.Traces, error) {

}

func  createLogsProcessor(
	ctx context.Context,
	set processor.CreateSettings,
	c component.Config,
) (processor.Logs, error) {

}