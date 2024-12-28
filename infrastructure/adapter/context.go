package adapter

import (
	"sync"
	"time"
)

func NewContext() *Context {
	return &Context{
		Keys: make(map[string]any),
	}
}

type Context struct {
	Token  string
	Scheme string
	Keys   map[string]any
	mu     sync.RWMutex
}

func (c *Context) Get(key string) (value any, has bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, has = c.Keys[key]
	return
}

func (t *Context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (t *Context) Done() <-chan struct{} {
	return nil
}

func (t *Context) Err() error {
	return nil
}

func (t *Context) Value(key any) any {
	if k, ok := key.(string); ok {
		if v, has := t.Get(k); has {
			return v
		}
	}
	return nil
}
