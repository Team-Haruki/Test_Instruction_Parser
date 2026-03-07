package sekai

import (
	"Haruki-Command-Parser/internal/handler"
	"Haruki-Command-Parser/internal/parser"
)

func (sekaiHandlers) DeckEventHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/活动组卡", "/活动组队", "/活动卡组", "/活动配队", "/组卡", "/组队", "/配队",
			"/指定属性组卡", "/指定属性组队", "/指定属性卡组", "/指定属性配队",
			"/模拟组卡", "/模拟配队", "/模拟组队", "/模拟卡组",
			"/pjsk event card", "/pjsk event deck", "/pjsk deck",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleDeck, "deck-event"), nil
		},
	}
}

func (sekaiHandlers) DeckChallengeHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/挑战组卡", "/挑战组队", "/挑战卡组", "/挑战配队", "/pjsk challenge card", "/pjsk challenge deck",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleDeck, "deck-challenge"), nil
		},
	}
}

func (sekaiHandlers) DeckNoEventHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/长草组卡", "/长草组队", "/长草卡组", "/长草配队", "/最强卡组", "/最强组卡", "/最强组队", "/最强配队",
			"/pjsk no event deck", "/pjsk best deck",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleDeck, "deck-no-event"), nil
		},
	}
}

func (sekaiHandlers) DeckBonusHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/加成组卡", "/加成组队", "/加成卡组", "/加成配队", "/控分组卡", "/控分组队", "/控分卡组", "/控分配队",
			"/pjsk bonus deck", "/pjsk bonus card",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleDeck, "deck-bonus"), nil
		},
	}
}

func (sekaiHandlers) DeckMysekaiHandle() handler.SekaiCommandHandlerConfig {
	return handler.SekaiCommandHandlerConfig{
		Commands: []string{
			"/烤森组卡", "/烤森组队", "/烤森卡组", "/烤森配队", "/ms组卡", "/ms组队", "/ms卡组", "/ms配队",
			"/mysekai deck", "/pjsk mysekai deck",
		},
		HandleFunc: func(ctx handler.SekaiHandlerContext) (interface{}, error) {
			return makeResolvedCmd(ctx, parser.ModuleDeck, "deck-mysekai"), nil
		},
	}
}
