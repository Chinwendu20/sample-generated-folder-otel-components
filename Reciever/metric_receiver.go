package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
)
// You can change struct name
type projectMetricReceiver struct{


}

func (projectReceiver *projectMetricReceiver)Start(ctx context.Context, host component.Host)error{
	return  nil

}

func (projectReceiver *projectMetricReceiver)Shutdown(ctx context.Context)error{
	return nil
}




