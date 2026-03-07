package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) CardDetailHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/card-detail", "/卡面", "/详情", "/查卡",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleCard, "card-detail"), nil
		},
	}
}

func (sekaiHandlers) CardListHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/查牌", "/查卡牌", "/卡牌列表", "/card", "/cards", "/pjsk card", "/pjsk member",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleCard, "card-list"), nil
		},
	}
}

func (sekaiHandlers) CardBoxHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/查箱", "/查框", "/卡牌一览", "/卡面一览", "/卡一览", "/box", "/card-box", "/pjsk box",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleCard, "card-box"), nil
		},
	}
}
