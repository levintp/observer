// Package to implement a simple and fast FIFO queue.
package queue

import (
	"errors"
)

type Queue[T comparable] []T

func (queue *Queue[T]) Pop() (T, error) {
	var top T
	var err error

	// Check if the queue is empty
	if len(*queue) == 0 {
		err = errors.New("queue is empty")
	} else {
		// Dequeue the top element from the queue and shrink the queue.
		top = (*queue)[0]
		*queue = (*queue)[1:]
	}

	return top, err
}

func (queue *Queue[T]) Push(item T) {
	*queue = append(*queue, item)
}
