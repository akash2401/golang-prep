package main

import (
	"fmt"
	"sync"
)

// Stack - follows LIFO principle - last in first out - last element to enter the stack is evicted the first
// operations - Push(Add an element to stack), Pop(Delete an element from stack), Peek(Returns top most element of stack)
// Special conditions - Underflow(When stack is empty and we try to delete an element) and Overflow(When stack is full and we try to insert an element - Not implemented in this example)
// This implementation is go-routine safe as well and that is because we are making use of mutex

type Stack struct {
	stackArr []int
	sMux     sync.Mutex
}

func (s *Stack) Push(ele int) {
	s.stackArr = append(s.stackArr, ele)
}

func (s *Stack) Pop() {
	tempArr := s.stackArr[:len(s.stackArr)-1]
	s.stackArr = tempArr

}

func (s *Stack) Peek() int {
	return s.stackArr[len(s.stackArr)-1]
}

func main() {
	// example to test out functionality
	s := Stack{}
	s.Push(3)
	fmt.Println(s.stackArr)
	s.Push(5)
	fmt.Println(s.stackArr)
	s.Push(9)
	fmt.Println(s.stackArr)
	s.Pop()
	fmt.Println("Elements inside stack: ", s.stackArr)
	fmt.Println("Element at top: ", s.Peek())
}
