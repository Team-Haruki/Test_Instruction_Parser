package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) SKLineHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-line", "/sk线", "/榜线", "/pjsk sk line", "/pjsk board line", "/skl",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-line"), nil
		},
	}
}

func (sekaiHandlers) SKQueryHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-query", "/sk查询", "/sk查分", "/pjsk sk board", "/pjsk board",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-query"), nil
		},
	}
}

func (sekaiHandlers) SKCheckRoomHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-check-room", "/sk查房", "/查房", "/cf", "/pjsk查房", "/csb", "/冲水板", "/pjsk冲水板",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-check-room"), nil
		},
	}
}

func (sekaiHandlers) SKSpeedHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-speed", "/sk时速", "/时速线", "/pjsk sk speed", "/pjsk board speed", "/sks", "/skv", "/sktime",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-speed"), nil
		},
	}
}

func (sekaiHandlers) SKPlayerTraceHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-player-trace", "/sk玩家轨迹", "/玩家轨迹", "/ptr", "/pjsk玩家追踪", "/pjsk ptr",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-player-trace"), nil
		},
	}
}

func (sekaiHandlers) SKRankTraceHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-rank-trace", "/sk档线轨迹", "/档线轨迹", "/rtr", "/skt", "/sklt", "/sktl", "/pjsk追踪", "/pjsk sk追踪",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-rank-trace"), nil
		},
	}
}

func (sekaiHandlers) SKWinrateHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk-winrate", "/sk胜率", "/胜率预测", "/pjsk winrate predict", "/5v5预测", "/5v5胜率",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleSK, "sk-winrate"), nil
		},
	}
}
