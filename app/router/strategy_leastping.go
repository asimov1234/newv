package router

import (
	"context"

	"github.com/asimov/newv/common/dice"
)

type LeastPingStrategy struct {
	ctx context.Context
}

func (l *LeastPingStrategy) GetPrincipleTarget(strings []string) []string {
	return []string{l.PickOutbound(strings)}
}

func (l *LeastPingStrategy) InjectContext(ctx context.Context) {
	l.ctx = ctx
	// Observatory functionality removed
}

func (l *LeastPingStrategy) PickOutbound(candidates []string) string {
	// Observatory functionality removed - using random selection as fallback
	count := len(candidates)
	if count == 0 {
		return ""
	}
	return candidates[dice.Roll(count)]
}
