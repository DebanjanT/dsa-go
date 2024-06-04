// Queue is First In First Out linear Data Structure
// Almost same as stack but not LIFO. Its FIFO
package main

import "fmt"

type queue struct {
	data []int
}

func (q *queue) Enqueue(v int) {
	q.data = append(q.data, v)
}

func (q *queue) Dequeue() int {
	dequeuedElement := q.data[0]
	q.data = q.data[1:]
	return dequeuedElement
}

func GoQueue() {
	Q := queue{}
	Q.Enqueue(1)
	Q.Enqueue(2)
	Q.Enqueue(3)
	Q.Enqueue(4)

	fmt.Println(Q.data)

	Q.Dequeue()
	Q.Dequeue()
	fmt.Println(Q.data)
}
