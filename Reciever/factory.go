
package project

import (
	"context"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/component"


)

const (
	// typeStr is the type of the receiver
	typeStr = "project"
)


// NewFactory creates a receiver factory
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the receiver type that you would like, change the second parameter as you like
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		// receiver.WithMetrics(createMetricsReceiver, component.StabilityLevelBeta),
		// receiver.WithTraces(createTracesReceiver, component.StabilityLevelBeta),
		// receiver.WithLogs(createLogsReceiver, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &Config{

	}
}

// createMetricsReceiver creates a metrics receiver based on this config.
func createMetricsReceiver(
	ctx context.Context,
	set receiver.CreateSettings,
	c component.Config,
) (receiver.Metrics, error) {

	
}

// createTracesReceiver creates a trace receiver based on this config.
func  createTracesReceiver(
	ctx context.Context,
	set receiver.CreateSettings,
	c component.Config,
) (receiver.Traces, error) {

}

func  createLogsReceiver(
	ctx context.Context,
	set receiver.CreateSettings,
	c component.Config,
) (receiver.Logs, error) {

}