package handler

import (
	"Haruki-Command-Parser/internal/parser"
	sekairegion "Haruki-Command-Parser/internal/sekai_region"
	"fmt"
	"slices"
	"strings"
)

type SekaiHandlerContext struct {
	HandlerContext
	region             *sekairegion.SekaiRegion // 区服
	originalTriggerCmd string                   // 原始触发命令，未去除区服前缀
	prefixArg          string                   // 额外前缀
	uidArg             string                   // UID参数 /u / uid / @
	flags              map[string]bool          // -verbose, -preview, -help 等开关
}

func (s *SekaiHandlerContext) Region() *sekairegion.SekaiRegion {
	return s.region
}
func (s *SekaiHandlerContext) PrefixArg() string {
	return s.prefixArg
}
func (s *SekaiHandlerContext) Flags() map[string]bool {
	return s.flags
}
func (s *SekaiHandlerContext) SetArgs(args string) {
	s.argText = args
}

type SekaiCommandHandler struct {
	BaseCommandHandler
	regions    []*sekairegion.SekaiRegion
	prefixArgs []string
	handleFunc func(SekaiHandlerContext) (interface{}, error)
}

func (skh *SekaiCommandHandler) Handle(ctx Context) (interface{}, error) {
	if skh.handleFunc == nil {
		cmdName := "未定义"
		if len(skh.commands) > 0 {
			cmdName = skh.commands[0]
		}
		return nil, fmt.Errorf("Sekai 命令处理器 %s 没有处理方法", cmdName)
	}
	// 处理指令区服前缀
	var cmdRegion *sekairegion.SekaiRegion
	originalTriggerCmd := ctx.GetTriggerCmd()
	triggerCmd := originalTriggerCmd
	for _, region := range skh.regions {
		cmdRegionPrefix := fmt.Sprintf("/%s", region.Id())
		if strings.HasPrefix(triggerCmd, cmdRegionPrefix) {
			cmdRegion = region
			triggerCmd = strings.Replace(triggerCmd, cmdRegionPrefix, "/", 1)
			break
		}
	}
	// 处理前缀参数
	prefixArg := ""
	for _, prefix := range skh.prefixArgs {
		cmdPrefix := fmt.Sprintf("/%s", prefix)
		if strings.HasPrefix(triggerCmd, cmdPrefix) {
			prefixArg = prefix
			triggerCmd = strings.Replace(triggerCmd, cmdPrefix, "/", 1)
			break
		}
	}
	// TODO: 如果没有指定区服，并且用户有默认区服，并且用户默认区服在可用区服列表中，则使用用户的默认区服

	// 如果没有指定区服，并且用户没有默认区服，则使用指令的默认区服
	if cmdRegion == nil && len(skh.regions) > 0 {
		cmdRegion = skh.regions[0]
	}
	// TODO: 测试时处理账号参数等
	args := ctx.GetArgs()

	// 提取通用 flags 和 region flag (不需要 nicknames 也能提取这些)
	ext := parser.NewExtractor(nil)
	flags := make(map[string]bool)

	regRes := ext.ExtractRegion(args)
	if regRes.Value != "" {
		if r := sekairegion.GetRegionById(regRes.Value); r != nil {
			cmdRegion = r
		}
	}
	args = regRes.Remaining

	verbRes := ext.ExtractVerbose(args)
	flags["is_verbose"] = verbRes.Value
	args = verbRes.Remaining

	preRes := ext.ExtractPreview(args)
	flags["is_preview"] = preRes.Value
	args = preRes.Remaining

	helpRes := ext.ExtractHelp(args)
	flags["is_help"] = helpRes.Value
	args = helpRes.Remaining

	skCtx := &SekaiHandlerContext{
		HandlerContext: HandlerContext{
			Context:     ctx,
			triggerCmd:  triggerCmd,
			argText:     args,
			messageType: ctx.GetMessageType(),
			message:     ctx.GetMessage(),
			event:       ctx.GetEvent(),
			messageId:   ctx.GetMessageId(),
			userId:      ctx.GetUserId(),
			senderName:  ctx.GetSenderName(),
			groupId:     ctx.GetGroupId(),
		},
		region:             cmdRegion,
		originalTriggerCmd: originalTriggerCmd,
		prefixArg:          prefixArg,
		uidArg:             "",
	}
	skCtx.region = cmdRegion
	skCtx.originalTriggerCmd = originalTriggerCmd
	skCtx.prefixArg = prefixArg
	skCtx.argText = args
	skCtx.flags = flags
	return skh.handleFunc(*skCtx)
}

// 工具结构体，用于初始化SekaiCommandHandler
type SekaiCommandHandlerConfig struct {
	Commands   []string
	Priority   int
	Helper     string
	Regions    []*sekairegion.SekaiRegion
	PrefixArgs []string
	HandleFunc func(SekaiHandlerContext) (interface{}, error)
}

var DefaultRegions = sekairegion.Regions

func (cfg SekaiCommandHandlerConfig) ToSekaiCommandHandler() *SekaiCommandHandler {
	// 注册包含区服前缀的所有指令
	if len(cfg.Regions) == 0 {
		cfg.Regions = DefaultRegions
	}
	if len(cfg.PrefixArgs) == 0 {
		cfg.PrefixArgs = []string{""}
	}
	allRegionCommands := make(map[string]bool, len(cfg.Commands)*len(cfg.Regions)*len(cfg.PrefixArgs))
	for _, prefix := range cfg.PrefixArgs {
		for _, region := range cfg.Regions {
			for _, cmd := range cfg.Commands {
				allRegionCommands[cmd] = true
				allRegionCommands[strings.Replace(cmd, "/", fmt.Sprintf("/%s", prefix), 1)] = true
				allRegionCommands[strings.Replace(cmd, "/", fmt.Sprintf("/%s%s", region.Id(), prefix), 1)] = true
			}
		}
	}
	// 去除重复指令
	cfg.Commands = []string{}
	for cmd := range allRegionCommands {
		cfg.Commands = append(cfg.Commands, cmd)
	}
	slices.Sort(cfg.Commands)
	// 默认优先级
	if cfg.Priority == 0 {
		cfg.Priority = 100
	}
	return &SekaiCommandHandler{
		BaseCommandHandler{
			cfg.Commands,
			cfg.Priority,
			cfg.Helper,
			nil,
		},
		cfg.Regions,
		cfg.PrefixArgs,
		cfg.HandleFunc,
	}
}
