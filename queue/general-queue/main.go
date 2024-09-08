package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	queue     []int
	queueLock sync.Mutex
	front     int
	rear      int
	size      int
}

func (q *Queue) Enqueue(ele int) {
	q.queueLock.Lock()
	defer q.queueLock.Unlock()
	if q.front == q.size {
		fmt.Println("Invalid enqueue operation performed. Queue is full.")
		return
	}
	q.queue[q.front] = ele
	q.front += 1
	fmt.Println("Element inserted in queue: ", ele, " at position: ", q.front)
}

func (q *Queue) Dequeue() {
	q.queueLock.Lock()
	defer q.queueLock.Unlock()
	if q.front == q.size {
		fmt.Println("Invalid dequeue operation performed. Queue is empty.")
		return
	}
	eleToBeDeleted := q.queue[q.rear]
	q.queue[q.rear] = -1
	q.rear += 1
	fmt.Println("Element deleted from queue: ", eleToBeDeleted, " at position: ", q.front)
}

func (q *Queue) InitializeQueue(size int) {
	q.size = size
	q.front = 0
	q.rear = 0
	q.queue = make([]int, q.size)
}

func main() {
	q := Queue{}
	q.InitializeQueue(4)
	// fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])
	q.Enqueue(5)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])
	q.Enqueue(11)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])

	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])
	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])

	q.Enqueue(13)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])
	q.Enqueue(17)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])
	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, ": value at front : ", q.queue[q.front-1], " - rear val: ", q.rear, " : value at rear : ", q.queue[q.rear])
	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue)
	q.Enqueue(13)
	q.Dequeue()
}
