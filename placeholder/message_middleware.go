package placeholder

import (
	"errors"
	"io"
	"webPlatform/config"
	"webPlatform/pipeline"
	"webPlatform/services"
)

type SimpleMessageComponent struct{}

func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext)) {
	var cfg config.Configuration
	services.GetService(&cfg)
	msg, ok := cfg.GetString("main:message")
	if ok {
		io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("Cannot find config setting"))
	}
	next(ctx)
}
