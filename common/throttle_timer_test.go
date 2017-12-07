package common

import (
	"sync"
	"testing"
	"time"

	// make govet noshadow happy...
	asrt "github.com/stretchr/testify/assert"
)

type thCounter struct {
	input <-chan struct{}
	mtx   sync.Mutex
	count int
}

func (c *thCounter) Increment() {
	c.mtx.Lock()
	c.count++
	c.mtx.Unlock()
}

func (c *thCounter) Count() int {
	c.mtx.Lock()
	val := c.count
	c.mtx.Unlock()
	return val
}

// Read should run in a go-routine and
// updates count by one every time a packet comes in
func (c *thCounter) Read() {
	// note, since this channel never closes, this will never end
	// if thCounter was used in anything beyond trivial test cases.
	// it would have to be smarter.
	for range c.input {
		c.Increment()
	}
}

func TestThrottle(test *testing.T) {
	assert := asrt.New(test)

	ms := 50
	delay := time.Duration(ms) * time.Millisecond
	shortwait := time.Duration(ms/2) * time.Millisecond
	longwait := time.Duration(2) * delay
	t := NewThrottleTimer("foo", delay)

	// start at 0
	c := &thCounter{input: t.Ch}
	assert.Equal(0, c.Count())
	go c.Read()

	// waiting does nothing
	time.Sleep(longwait)
	assert.Equal(0, c.Count())

	// send one event adds one
	t.Set()
	time.Sleep(longwait)
	assert.Equal(1, c.Count())

	// send a burst adds one
	for i := 0; i < 5; i++ {
		t.Set()
	}
	time.Sleep(longwait)
	assert.Equal(2, c.Count())

	// keep cancelling before it is ready
	for i := 0; i < 10; i++ {
		t.Set()
		time.Sleep(shortwait)
		t.Unset()
	}
	time.Sleep(longwait)
	assert.Equal(2, c.Count())

	// a few unsets do nothing...
	for i := 0; i < 5; i++ {
		t.Unset()
	}
	assert.Equal(2, c.Count())

	// send 12, over 2 delay sections, adds 3
	short := time.Duration(ms/5) * time.Millisecond
	for i := 0; i < 13; i++ {
		t.Set()
		time.Sleep(short)
	}
	time.Sleep(longwait)
	assert.Equal(5, c.Count())

	stopped := t.Stop()
	assert.True(stopped)
}