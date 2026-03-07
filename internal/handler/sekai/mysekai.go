package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) MysekaiResourceHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/mysekai-resource", "/mysekai资源", "/烤森资源", "/msr", "/msmap", "/msa",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMysekai, "mysekai-resource"), nil
		},
	}
}

func (sekaiHandlers) MysekaiFixtureListHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/mysekai-fixture-list", "/mysekai家具列表", "/烤森家具列表", "/msf",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMysekai, "mysekai-fixture-list"), nil
		},
	}
}

func (sekaiHandlers) MysekaiFixtureDetailHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/mysekai-fixture-detail", "/mysekai家具详情", "/烤森家具详情",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMysekai, "mysekai-fixture-detail"), nil
		},
	}
}

func (sekaiHandlers) MysekaiDoorUpgradeHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/mysekai-door-upgrade", "/mysekai大门升级", "/烤森大门升级", "/msg", "/msgate",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMysekai, "mysekai-door-upgrade"), nil
		},
	}
}

func (sekaiHandlers) MysekaiMusicRecordHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/mysekai-music-record", "/mysekai唱片", "/烤森唱片", "/msm", "/mss",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMysekai, "mysekai-music-record"), nil
		},
	}
}

func (sekaiHandlers) MysekaiTalkListHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/mysekai-talk-list", "/mysekai对话列表", "/烤森对话列表",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMysekai, "mysekai-talk-list"), nil
		},
	}
}
