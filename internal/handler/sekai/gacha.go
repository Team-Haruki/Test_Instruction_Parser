package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) GachaHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/卡池", "/查卡池", "/卡池列表", "/卡池一览", "/抽卡", "/gacha", "/gacha-list", "/pjsk gacha",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleGacha, "gacha"), nil
		},
	}
}
