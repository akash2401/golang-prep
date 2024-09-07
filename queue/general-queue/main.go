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
	}
	if q.front == -1 {
		q.rear += 1
	}
	q.front += 1
	q.queue[q.front] = ele
	fmt.Println("Element inserted in queue: ", ele, " at position: ", q.front)
}

func (q *Queue) Dequeue() {
	q.queueLock.Lock()
	defer q.queueLock.Unlock()
	q.rear += 1
	eleToBeDeleted := q.queue[0]
	tempArr := make([]int, q.size)
	copy(tempArr, q.queue[1:])
	tempArr[len(tempArr)-1] = -1
	q.queue = tempArr
	fmt.Println("Element deleted from queue: ", eleToBeDeleted, " at position: ", q.front)
}

func (q *Queue) InitializeQueue(size int) {
	q.size = size
	q.front = -1
	q.rear = -1
	q.queue = make([]int, q.size)
}

func main() {
	q := Queue{}
	q.InitializeQueue(4)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)
	q.Enqueue(5)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)
	q.Enqueue(11)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)

	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)
	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)

	q.Enqueue(13)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)
	q.Enqueue(17)
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)
	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)
	q.Dequeue()
	fmt.Println("Printing queue elements: ", q.queue, " - front val: ", q.front, " - rear val: ", q.rear)

}
