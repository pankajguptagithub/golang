package main

import (
	"fmt"
)

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}
func (q *Queue) Traverse() {
	if q.IsEmpty() == false {
		fmt.Println("Items of the queue are")
		for i := 0; i < len(q.items); i++ {
			fmt.Println(q.items[i])
		}
	}
}
func (q *Queue) Dequeue() {
	if q.IsEmpty() == false {
		item := q.items[0]
		fmt.Println("Item deleted from the queue is ", item)
		q.items = q.items[1:]
	}
}
func (q *Queue) IsEmpty() bool {
	if q.Size() == 0 {
		return true
	} else {
		return false
	}
	//other implementation
	//return len(q.items) == 0
}
func (q *Queue) Size() int {
	return len(q.items)
}
func main() {
	var q Queue
	for i := 10; i <= 20; i++ {
		q.Enqueue(i)
	}
	fmt.Println("Size of the queue is ", q.Size())
	q.Traverse()
	q.Dequeue()
	q.Dequeue()
	q.Traverse()
	fmt.Println("Size of the queue is ", q.Size())
}
