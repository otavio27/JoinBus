/*
 MODULE: TTKQueue.go
 AUTHOR: Leo Schneider <schleo@outlook.com>
 DATE  : 3 June 2018
 INFO  : This module implements a FIFO structure and its related methods
*/

package ttktools

type ttkqueue struct {
	queue []interface{}
}

// Queue Creates a new queue
func Queue() *ttkqueue {
	t := ttkqueue{}
	return &t
}

// Push - puts a new item in the queue
func (t *ttkqueue) Push(item interface{}) {
	t.queue = append(t.queue, item)
}

// Pop - returns the first item from the qeue
func (t *ttkqueue) Pop() interface{} {
	if t.Length() > 0 {
		ret := t.queue[0]
		t.queue = t.queue[1:]
		return ret
	}
	return nil
}

// Rpop - put the first item back in the queue
func (t *ttkqueue) Rpop() {
	t.Push(t.Pop())
}

// Length - returns the number of items in the queue
func (t *ttkqueue) Length() int {
	return len(t.queue)
}
