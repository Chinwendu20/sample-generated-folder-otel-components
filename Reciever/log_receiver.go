package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
)
// You can change struct name
type projectLogReceiver struct{


}

func (projectReceiver *projectLogReceiver)Start(ctx context.Context, host component.Host)error{
	return  nil

}

func (projectReceiver *projectLogReceiver)Shutdown(ctx context.Context)error{
	return nil
}




