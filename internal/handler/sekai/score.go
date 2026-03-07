package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) ScoreControlHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/分数", "/查分数", "/pjsk score", "/score control",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleScore, "score-control"), nil
		},
	}
}

func (sekaiHandlers) ScoreCustomRoomHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/自定义房间分数", "/自定义分数", "/custom room score", "/pjsk custom room score",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleScore, "score-custom-room"), nil
		},
	}
}

func (sekaiHandlers) ScoreMusicMetaHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/曲目meta", "/music meta", "/pjsk music meta",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleScore, "score-music-meta"), nil
		},
	}
}

func (sekaiHandlers) ScoreMusicBoardHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/曲目榜", "/歌曲比较", "/music board", "/pjsk music board",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleScore, "score-music-board"), nil
		},
	}
}
