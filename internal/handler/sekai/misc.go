package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) MiscBirthdayHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/角色生日", "/生日贺图", "/pjsk chara birthday", "/chara birthday",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleMisc, "misc-birthday"), nil
		},
	}
}

func (sekaiHandlers) HelpHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/help", "/帮助",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleHelp, "help"), nil
		},
	}
}

func (sekaiHandlers) ProfileHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/sk", "/个人中心", "/个人信息", "/名片", "/pjsk profile", "/profile",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleProfile, "profile"), nil
		},
	}
}
