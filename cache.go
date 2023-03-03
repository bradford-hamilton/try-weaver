package main

import (
	"context"
	"sync"
)

/*
	Semantics:
		When implementing a component, there are three semantic details to keep in mind:
			1. A component's state is not persisted.
			2. A component's methods may be invoked concurrently.
			3. There may be multiple replicas of a component.
		A Cache's state is not persisted, so if a Cache replica fails, its data is lost.
		Any state that needs to be persisted should be persisted explicitly. A Cache's
		methods may be invoked concurrently, so it's essential that we guard access to
		data with the mutex. There may be multiple replicas of a Cache component, so it
		is not guaranteed that one client's Get will be routed to the same replica as
		another client's Put. For this example, this means that the Cache has weak
		consistency.
*/

type Cache interface {
	Put(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type cache struct {
	mu   sync.Mutex
	data map[string]string
}

func (c *cache) Put(_ context.Context, key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
	return nil
}

func (c *cache) Get(_ context.Context, key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[key], nil
}
