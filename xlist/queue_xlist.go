// Package xlist is a list data struct expand package.
package xlist

import (
	"container/list"
	"errors"
	"sync"
)

// Queue implement a thread-safe queue.
type Queue struct {
	li *list.List
	// if li == cap, Push will return error ErrQueueFull. It's read only.
	cap int
	mu  sync.RWMutex
}

// ErrQueueFull is a full queue error.
// If it returns, the Push call will fail and discard the push element.
var (
	ErrQueueFull = errors.New("push in a full li")
)

// NewQueue ...
func NewQueue(cap int) (*Queue, error) {
	q := new(Queue)

	q.li = list.New()
	q.cap = cap

	return q, nil
}

// Cap returns cap maximum cap of queue.
func (q *Queue) Cap() int {
	return q.cap
}

// Push e into q.
func (q *Queue) Push(e interface{}) error {
	if q.li.Len() >= q.cap {
		return ErrQueueFull
	}

	q.mu.Lock()
	q.li.PushBack(e)
	q.mu.Unlock()

	return nil
}

// PushFront ...
func (q *Queue) PushFront(e interface{}) error {
	if q.li.Len() >= q.cap {
		return ErrQueueFull
	}

	q.mu.Lock()
	q.li.PushFront(e)
	q.mu.Unlock()

	return nil
}

// PopFront ...
func (q *Queue) PopFront(l int, f func(e interface{})) (int, error) {
	if l > q.Len() {
		l = q.Len()
	}

	for i := 0; i < l; i++ {
		q.mu.Lock()
		e := q.li.Front()
		if e != nil {
			q.li.Remove(e)
		}
		q.mu.Unlock()

		if e != nil {
			f(e.Value)
		}
	}

	return l, nil
}

// Len ...
func (q *Queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.len()
}

func (q *Queue) len() int {
	return q.li.Len()
}

// Clear ...
func (q *Queue) Clear() error {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.li.Init()

	return nil
}

// Dump will dump all queue values.
func (q *Queue) Dump(l int) []interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	res := make([]interface{}, 0, l)

	if l >= q.len() || l == 0 {
		l = q.len()
	}
	e := q.li.Front()
	for i := 0; i < l && e != nil; i++ {
		res = append(res, e.Value)
		e = e.Next()
	}

	return res
}
