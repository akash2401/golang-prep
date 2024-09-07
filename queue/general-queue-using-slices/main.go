package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	queue     []int
	queueLock sync.Mutex
}

func (q *Queue) Enqueue(ele int) {

	q.queueLock.Lock()
	defer q.queueLock.Unlock()
	q.queue = append(q.queue, ele)
}

func (q *Queue) Dequeue() {

	q.queueLock.Lock()
	defer q.queueLock.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("Empty queue.")
		return
	}
	q.queue = q.queue[1:]
}

func (q *Queue) Peek() int {

	q.queueLock.Lock()
	defer q.queueLock.Unlock()
	return q.queue[0]
}

func main() {

	q := Queue{}
	q.Enqueue(5)
	fmt.Println("Elements inside queue: ", q.queue)
	q.Enqueue(9)
	fmt.Println("Elements inside queue: ", q.queue)
	q.Dequeue()
	fmt.Println("Elements inside queue: ", q.queue)
	fmt.Println(q.Peek())
	q.Enqueue(11)
	fmt.Println("Elements inside queue: ", q.queue)
	q.Dequeue()
	fmt.Println("Elements inside queue: ", q.queue)
	fmt.Println(q.Peek())
	q.Enqueue(13)
	fmt.Println("Elements inside queue: ", q.queue)
	fmt.Println(q.Peek())
	q.Dequeue()
	fmt.Println("Elements inside queue: ", q.queue)
	q.Dequeue()
	fmt.Println("Elements inside queue: ", q.queue)
	q.Dequeue()
	fmt.Println("Elements inside queue: ", q.queue)
}
