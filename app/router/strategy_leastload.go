package router

import (
	"context"

	"github.com/xtls/xray-core/common/dice"
)

// LeastLoadStrategy represents a least load balancing strategy
// Observatory functionality has been removed, fallback to random selection
type LeastLoadStrategy struct {
	settings *StrategyLeastLoadConfig
	ctx      context.Context
}

func (l *LeastLoadStrategy) GetPrincipleTarget(strings []string) []string {
	return strings
}

func (l *LeastLoadStrategy) InjectContext(ctx context.Context) {
	l.ctx = ctx
	// Observatory functionality removed
}

func (l *LeastLoadStrategy) PickOutbound(candidates []string) string {
	// Observatory functionality removed - using random selection as fallback
	count := len(candidates)
	if count == 0 {
		return ""
	}
	return candidates[dice.Roll(count)]
}
