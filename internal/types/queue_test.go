package types_test

import (
	"testing"

	"github.com/levintp/observer/internal/types"
)

func TestPush(t *testing.T) {
	queue := make(types.Queue[int], 0)

	// Push some values into the queue.
	queue.Push(123)
	queue.Push(456)
	queue.Push(789)

	// Check length
	if len(queue) != 3 {
		t.Error("queue size has not changed")
	}

	// Check values.
	if queue[1] != 456 {
		t.Error("pushed item not in queue")
	}
}

func TestPop(t *testing.T) {
	queue := make(types.Queue[int], 0)

	// Assert error is thrown when queue is empty.
	if _, err := queue.Pop(); err == nil {
		t.Error("no error thrown after popping an empty queue")
	}

	// Manually insert some values into the queue slice.
	queue = append(queue, 123)
	queue = append(queue, 456)
	queue = append(queue, 789)

	// Assert pop functionality is correct.
	top, err := queue.Pop()
	if err != nil {
		t.Errorf("error thrown when popping queue: %v", err)
	}
	if top != 123 {
		t.Error("wrong value popped from queue")
	}

	// Assert queue length has decreased.
	if len(queue) != 2 {
		t.Error("queue's length hasn't changed after pop")
	}
}
