package safecount

import "sync"

// Counter counts
type Counter struct {
	mu    sync.Mutex
	value int
}

// Inc increments a counter's value
func (c *Counter) Inc() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// Value returns the counter's value
func (c *Counter) Value() int {
	return c.value
}
