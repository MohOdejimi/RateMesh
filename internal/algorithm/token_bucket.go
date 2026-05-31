package algorithm

import (
	"context"
	"sync"
	"time"

	"github.com/MohOdejimi/RateMesh/internal/rule"
)

type bucketState struct {
	tokens float64 
	lastRefill time.Time 
	mu sync.Mutex
}

type TokenBucket struct {
	bucket map[string]*bucketState
	mu sync.Mutex
}

func NewTokenBucket() *TokenBucket {
	return &TokenBucket{
		bucket: make(map[string]*bucketState),
	}
}

func (tb *TokenBucket) Allow(_ context.Context, clientID string, rule rule.Rule) (Decision, error) {
	tb.mu.Lock()
	state, exists := tb.bucket[clientID]
	if !exists {
		tb.bucket[clientID] = &bucketState{
			tokens: float64(rule.Limit),
			lastRefill: time.Now(),
		}
		state = tb.bucket[clientID]
	}				
	tb.mu.Unlock()

	state.mu.Lock()
	passedTime := time.Since(state.lastRefill)
	refillRate := float64(rule.Limit) / rule.Window.Seconds()
	state.tokens += passedTime.Seconds() * refillRate
	if state.tokens > float64(rule.Limit) {
		state.tokens = float64(rule.Limit)
	}
	state.lastRefill = time.Now()		
	allowed := state.tokens >= 1
	if allowed {
		state.tokens -= 1
	}		
	remaining := int(state.tokens)
	reset := state.lastRefill.Add(rule.Window)			
	state.mu.Unlock()
	decision := Decision{
		Allowed: allowed,
		Remaining: remaining,		
		Limit: rule.Limit,
		Reset: reset,
	}			
	if !allowed {
		decision.RetryAfter = time.Duration((1 / refillRate) * float64(time.Second))
	}			
	return decision, nil
}