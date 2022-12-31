package project

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
)
// You can change struct name
type projectExtension struct{


}

func (projectExtension *projectExtension)Start(ctx context.Context, host component.Host)error{

}

func (projectExtension *projectExtension)Shutdown(ctx context.Context)error{
	return nil
}



