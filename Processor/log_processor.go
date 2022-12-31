package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/ptrace"
)
// You can change struct name
type projectLogProcessor struct{


}

func (projectLogProcessor *projectLogProcessor)Start(ctx context.Context, host component.Host)error{

}

func (projectLogProcessor *projectLogProcessor)Shutdown(ctx context.Context)error{
	return nil
}


func (projectLogProcessor *projectLogProcessor) consumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {

}

func (projectLogProcessor *projectLogProcessor) Capabilities() consumer.Capabilities {

}

