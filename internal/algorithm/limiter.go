package algorithm

import (
	"context"
	"time"

	"github.com/MohOdejimi/RateMesh/internal/rule"
)

type Decision struct {
	Allowed bool
	Remaining int 
	Limit int 
	Reset time.Time
	RetryAfter time.Duration
}

type Limiter interface {
	Allow(ctx context.Context, clientID string, rule rule.Rule) (Decision, error)
}