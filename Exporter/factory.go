package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	// typeStr is the type of the exporter
	typeStr = "project"
)


// NewFactory creates a Datadog exporter factory
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		// Uncomment the exporter type that you would like, change the second parameter as you like
		// component.StabilityLevelUndefined
		// component.StabilityLevelUnmaintained
		// component.StabilityLevelDeprecated
		// component.StabilityLevelDevelopment
		// component.StabilityLevelAlpha
		// component.StabilityLevelBeta
		// component.StabilityLevelStable
		// exporter.WithMetrics(f.createMetricsExporter, component.StabilityLevelBeta),
		// exporter.WithTraces(f.createTracesExporter, component.StabilityLevelBeta),
		// exporter.WithLogs(f.createLogsExporter, component.StabilityLevelAlpha),
	)
}

func createDefaultConfig() component.Config {

	return &Config{

	}
}

// createMetricsExporter creates a metrics exporter based on this config.
func createMetricsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	c component.Config,
) (exporter.Metrics, error) {

	
}

// createTracesExporter creates a trace exporter based on this config.
func  createTracesExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	c component.Config,
) (exporter.Traces, error) {

}

func  createLogsExporter(
	ctx context.Context,
	set exporter.CreateSettings,
	c component.Config,
) (exporter.Logs, error) {

}