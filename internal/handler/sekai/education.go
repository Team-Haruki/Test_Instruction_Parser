package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) EducationChallengeHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/挑战赛", "/挑战信息", "/挑战赛信息", "/挑战一览", "/每日挑战", "/pjsk challenge info", "/challenge info",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEducation, "education-challenge"), nil
		},
	}
}

func (sekaiHandlers) EducationPowerHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/加成信息", "/角色加成", "/加成一览", "/pjsk power bonus info", "/power bonus",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEducation, "education-power"), nil
		},
	}
}

func (sekaiHandlers) EducationAreaHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/区域道具", "/区域道具材料", "/道具升级", "/pjsk area item", "/area item",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEducation, "education-area"), nil
		},
	}
}

func (sekaiHandlers) EducationBondsHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/羁绊", "/角色羁绊", "/羁绊等级", "/羁绊信息", "/牵绊", "/牵绊等级", "/pjsk bonds", "/pjsk bond",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEducation, "education-bonds"), nil
		},
	}
}

func (sekaiHandlers) EducationLeaderHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/加成统计", "/领队统计", "/角色领队", "/pjsk leader count",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleEducation, "education-leader"), nil
		},
	}
}
