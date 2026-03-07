package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) EventListHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/活动列表", "/查活动列表", "/活动一览", "/events", "/event-list",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEvent, "event-list"), nil
		},
	}
}

func (sekaiHandlers) EventDetailHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/活动", "/查活动", "/event",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEvent, "event-detail"), nil
		},
	}
}
