package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
)
// You can change struct name
type projectTraceExporter struct{


}


func (projectExporter *projectTraceExporter)Start(ctx context.Context, host component.Host)error{
	return nil

}

func (projectExporter *projectTraceExporter)Shutdown(ctx context.Context)error{
	return nil
}


func (projectExporter *projectTraceExporter) ConsumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {

	return nil

}

