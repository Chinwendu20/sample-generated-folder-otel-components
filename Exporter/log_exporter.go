package project

import (
	"context"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/ptrace"
)
// You can change struct name
type projectLogExporter struct{


}

func (projectExporter *projectLogExporter)Start(ctx context.Context, host component.Host)error{
return nil
}

func (projectExporter *projectLogExporter)Shutdown(ctx context.Context)error{

	return nil
}


func (projectExporter *projectLogExporter) ConsumeTraces(
	ctx context.Context,
	td ptrace.Traces,
) (err error) {

}

