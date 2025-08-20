package router

import (
	"context"

	"github.com/asimov1234/newv/common/dice"
)

// RandomStrategy represents a random balancing strategy
type RandomStrategy struct {
	FallbackTag string

	ctx context.Context
}

func (s *RandomStrategy) InjectContext(ctx context.Context) {
	s.ctx = ctx
	// Observatory functionality removed
}

func (s *RandomStrategy) GetPrincipleTarget(strings []string) []string {
	return strings
}

func (s *RandomStrategy) PickOutbound(candidates []string) string {
	// Observatory functionality removed - using simple random selection
	count := len(candidates)
	if count == 0 {
		// goes to fallbackTag
		return ""
	}
	return candidates[dice.Roll(count)]
}
