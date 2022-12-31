package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
)
// You can change struct name
type projectMetricProcessor struct{


}

func (projectMetricProcessor *projectMetricProcessor)Start(ctx context.Context, host component.Host)error{

}

func (projectMetricProcessor *projectMetricProcessor)Shutdown(ctx context.Context)error{
	return nil
}


func (projectMetricProcessor *projectMetricProcessor) consumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {

}

func (projectMetricProcessor *projectMetricProcessor) Capabilities() consumer.Capabilities {

}

