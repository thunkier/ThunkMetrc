package services

import (
	"context"
	"log/slog"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type RateLimiterConfig struct {
	Enabled                     bool
	MaxGetPerSecondPerFacility  float64
	MaxGetPerSecondIntegrator   float64
	MaxConcurrentGetPerFacility int
	MaxConcurrentGetIntegrator  int

	// Resilience
	MaxRetries       int
	InitialBackoff   time.Duration
	MaxBackoff       time.Duration
	BackoffExponent  float64
}

func DefaultConfig() RateLimiterConfig {
	return RateLimiterConfig{
		Enabled:                     false,
		MaxGetPerSecondPerFacility:  50,
		MaxGetPerSecondIntegrator:   150,
		MaxConcurrentGetPerFacility: 10,
		MaxConcurrentGetIntegrator:  30,
		MaxRetries:                  10, // Default to high retry count
		InitialBackoff:              1 * time.Second,
		MaxBackoff:                  60 * time.Second,
		BackoffExponent:             2.0,
	}
}

type FacilityLimiter struct {
	rate *rate.Limiter
	sem  chan struct{}
}

type RateLimiter interface {
	Execute(ctx context.Context, facility string, isGet bool, op func() (interface{}, error)) (interface{}, error)
}

type StandardRateLimiter struct {
	config         RateLimiterConfig
	integratorRate *rate.Limiter
	integratorSem  chan struct{}
	facilitiesMu   sync.Mutex
	facilities     map[string]*FacilityLimiter
	logger         *slog.Logger
}

func NewRateLimiter(config RateLimiterConfig, logger *slog.Logger) RateLimiter {
	return &StandardRateLimiter{
		config:         config,
		integratorRate: rate.NewLimiter(rate.Limit(config.MaxGetPerSecondIntegrator), int(config.MaxGetPerSecondIntegrator)),
		integratorSem:  make(chan struct{}, config.MaxConcurrentGetIntegrator),
		facilities:     make(map[string]*FacilityLimiter),
		logger:         logger,
	}
}

func (r *StandardRateLimiter) Execute(ctx context.Context, facility string, isGet bool, op func() (interface{}, error)) (interface{}, error) {
	if !r.config.Enabled || !isGet {
		return r.executeWithRetry(ctx, op)
	}

	if r.logger != nil {
		r.logger.DebugContext(ctx, "Acquiring Integrator Semaphore", "facility", facility)
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r.integratorSem <- struct{}{}:
		defer func() { <-r.integratorSem }()
	}

	if facility != "" {
		fl := r.getFacilityLimiter(facility)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case fl.sem <- struct{}{}:
			defer func() { <-fl.sem }()
		}
	}

	if r.logger != nil {
		r.logger.DebugContext(ctx, "Waiting for Rate Limit", "facility", facility)
	}
	if err := r.integratorRate.Wait(ctx); err != nil {
		return nil, err
	}
	if facility != "" {
		fl := r.getFacilityLimiter(facility)
		if err := fl.rate.Wait(ctx); err != nil {
			return nil, err
		}
	}

	return r.executeWithRetry(ctx, op)
}

func (r *StandardRateLimiter) executeWithRetry(ctx context.Context, op func() (interface{}, error)) (interface{}, error) {
	var err error
	var res interface{}

	for attempt := 0; attempt <= r.config.MaxRetries; attempt++ {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

		res, err = op()
		if err == nil {
			return res, nil
		}

		errMsg := err.Error()
		is429 := strings.Contains(errMsg, "429") || strings.Contains(errMsg, "Too Many Requests")
		is5xx := strings.Contains(errMsg, "500") || strings.Contains(errMsg, "502") || strings.Contains(errMsg, "503") || strings.Contains(errMsg, "504")

		if !is429 && !is5xx {
			return nil, err
		}

		if attempt == r.config.MaxRetries {
			if r.logger != nil {
				r.logger.ErrorContext(ctx, "Max Retries Reached", "error", err, "attempt", attempt)
			}
			return nil, err // Final failure
		}

		// Calculate Backoff
		wait := r.calculateBackoff(attempt, err)
		
		if r.logger != nil {
			r.logger.WarnContext(ctx, "Retrying Request", "attempt", attempt+1, "wait_ms", wait.Milliseconds(), "error", err)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(wait):
			continue
		}
	}
	return nil, err
}

func (r *StandardRateLimiter) calculateBackoff(attempt int, _ error) time.Duration {
	backoffFloat := float64(r.config.InitialBackoff) * math.Pow(r.config.BackoffExponent, float64(attempt))
	
	if backoffFloat > float64(r.config.MaxBackoff) {
		backoffFloat = float64(r.config.MaxBackoff)
	}

	jitter := (rand.Float64() * 0.2) * backoffFloat
	total := backoffFloat + jitter

	return time.Duration(total)
}

func (r *StandardRateLimiter) getFacilityLimiter(facility string) *FacilityLimiter {
	r.facilitiesMu.Lock()
	defer r.facilitiesMu.Unlock()

	if fl, ok := r.facilities[facility]; ok {
		return fl
	}

	fl := &FacilityLimiter{
		rate: rate.NewLimiter(rate.Limit(r.config.MaxGetPerSecondPerFacility), int(r.config.MaxGetPerSecondPerFacility)),
		sem:  make(chan struct{}, r.config.MaxConcurrentGetPerFacility),
	}
	r.facilities[facility] = fl
	return fl
}
