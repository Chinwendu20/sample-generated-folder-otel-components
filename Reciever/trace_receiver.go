package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
)
// You can change struct name
type projectTraceReceiver struct{


}

func (projectReceiver *projectTraceReceiver)Start(ctx context.Context, host component.Host)error{
	return  nil

}

func (projectReceiver *projectTraceReceiver)Shutdown(ctx context.Context)error{
	return nil
}




