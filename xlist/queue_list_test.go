// Package xlist
package xlist

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q, err := NewQueue(1)
	assert.Nil(t, err)

	assert.Equal(t, 1, q.Cap())

	except := 1
	err = q.Push(except)
	assert.Nil(t, err)

	_, err = q.PopFront(1, func(e interface{}) {
		assert.Equal(t, except, e)
	})
	assert.Nil(t, err)
}

func TestQueue_FullError(t *testing.T) {
	q, err := NewQueue(1)
	assert.Nil(t, err)

	err = q.Push(1)
	assert.Nil(t, err)
	err = q.Push(1)
	assert.Equal(t, ErrQueueFull, err)
}

func TestQueue_Clear(t *testing.T) {
	q, err := NewQueue(1)
	assert.Nil(t, err)

	err = q.Push(1)
	assert.Nil(t, err)
	err = q.Push(1)
	assert.Equal(t, ErrQueueFull, err)

	err = q.Clear()
	assert.Nil(t, err)

	err = q.Push(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Len())
}

func TestQueue_PopFront(t *testing.T) {
	q, err := NewQueue(1)
	assert.Nil(t, err)

	assert.Equal(t, 1, q.Cap())

	except := 1
	err = q.Push(except)
	assert.Nil(t, err)

	_, err = q.PopFront(10, func(e interface{}) {
		assert.Equal(t, except, e)
	})
	assert.Nil(t, err)
}

func TestQueue_PushFront(t *testing.T) {
	except := []int{1, 2, 3, 4, 5}
	reverse := []interface{}{5, 4, 3, 2, 1}

	q, err := NewQueue(len(except))
	assert.Nil(t, err)

	assert.Equal(t, len(except), q.Cap())

	for i := 0; i < len(except); i++ {
		e := except[i]
		err = q.PushFront(e)
		assert.Nil(t, err)
	}
	err = q.PushFront(100)
	assert.Equal(t, ErrQueueFull, err)

	res := q.Dump(len(except))
	assert.Equal(t, reverse, res)
}

func TestQueueThreadSafe(t *testing.T) {
	var wait sync.WaitGroup
	cap := 1000
	q, err := NewQueue(cap)
	assert.Nil(t, err)

	for i := 0; i < cap; i++ {
		go func() {
			q.Push(i)
		}()
	}

	for i := 0; i < cap; i++ {
		wait.Add(1)

		go func() {
			q.PopFront(1, func(e interface{}) {
				// discard e.
			})
			wait.Done()
		}()
	}
	wait.Wait()

	assert.Equal(t, 0, q.Len())
}
