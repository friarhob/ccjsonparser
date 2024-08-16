package adt

import (
	"container/list"
	"errors"
)

type Queue struct {
	list *list.List
}

// NewQueue creates a new Queue
func NewQueue() *Queue {
	return &Queue{list: list.New()}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(item interface{}) {
	q.list.PushBack(item)
}

// Dequeue removes and returns the item at the front of the queue
func (q *Queue) Dequeue() (interface{}, error) {
	first := q.list.Front()
	if first == nil {
		return nil, errors.New("dequeuing from empty list")
	}
	q.list.Remove(first)
	return first.Value, nil
}

func (q *Queue) Peek() (interface{}, error) {
	first := q.list.Front()
	if first == nil {
		return nil, errors.New("peeking an empty list")
	}
	return first.Value, nil
}

// Size returns the number of items in the queue
func (q *Queue) Size() int {
	return q.list.Len()
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
