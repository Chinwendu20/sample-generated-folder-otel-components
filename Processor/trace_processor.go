package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
)
// You can change struct name
type projectTraceProcessor struct{


}

func (projectTraceProcessor *projectTraceProcessor)Start(ctx context.Context, host component.Host)error{

}

func (projectTraceProcessor *projectTraceProcessor)Shutdown(ctx context.Context)error{
	return nil
}


func (projectTraceProcessor *projectTraceProcessor) consumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {

}

func (projectTraceProcessor *projectTraceProcessor) Capabilities() consumer.Capabilities {

}

