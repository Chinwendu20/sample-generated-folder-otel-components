package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
)
// You can change struct name
type projectMetricExporter struct{


}

func (projectExporter *projectMetricExporter)Start(ctx context.Context, host component.Host)error{

}

func (projectExporter *projectMetricExporter)Shutdown(ctx context.Context)error{
	return nil
}


func (projectExporter *projectMetricExporter) ConsumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {

}

