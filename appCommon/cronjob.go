package appCommon

import (
	"context"
	goservice "github.com/lequocbinh04/go-sdk"
)

type handler func(ctx context.Context, serviceCtx goservice.ServiceContext) error

type Job struct {
	Name    string
	Cycle   string
	Handler handler
}
