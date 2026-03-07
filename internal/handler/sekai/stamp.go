package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) StampHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/贴纸", "/查贴纸", "/pjsk贴纸", "/pjsk stamp", "/pjsk bq", "/stamp",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleStamp, "stamp-list"), nil
		},
	}
}
