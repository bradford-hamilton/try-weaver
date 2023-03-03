package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

type reverser struct {
	weaver.Implements[Reverser]
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
	// Weaver logging from a component example
	// You can also add metrics, labels, tracing, profiling, etc
	logger := r.Logger()
	logger.Debug("A debug log.")
	logger.Info("An info log.")
	logger.Error("An error log.", fmt.Errorf("an error"))

	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}
