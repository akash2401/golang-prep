package main

import (
	"fmt"
	"sync"
)

type StackUsingArr struct {
	stackArr []int
	sMux     sync.Mutex
	top      int
	size     int
}

func (s *StackUsingArr) Push(ele int) {
	s.sMux.Lock()
	defer s.sMux.Unlock()
	if s.top == s.size {
		fmt.Println("Overflow")
		return
	}
	s.stackArr[s.top] = ele
	s.top++
}

func (s *StackUsingArr) Pop() {
	s.sMux.Lock()
	defer s.sMux.Unlock()
	s.top--
	if s.top < 0 {
		fmt.Println("Underflow")
		return
	}
	s.stackArr[s.top] = -1
	fmt.Println("pop top : ", s.top)
}

func (s *StackUsingArr) Peek() *int {
	s.sMux.Lock()
	defer s.sMux.Unlock()
	if s.top <= 0 {
		fmt.Println("no element present in the stack")
		return nil
	}
	return &s.stackArr[s.top-1]
}

func (s *StackUsingArr) InitializeStack(size int) {
	s.sMux.Lock()
	defer s.sMux.Unlock()
	s.size = size
	s.stackArr = make([]int, s.size)
	s.top = 0
}

func main() {
	// example 1 to test out functionality
	s := StackUsingArr{}
	s.InitializeStack(2)
	s.Push(3)
	fmt.Println(s.stackArr)
	s.Push(5)
	fmt.Println(s.stackArr)
	s.Push(9)
	fmt.Println(s.stackArr)
	s.Push(10)
	fmt.Println(s.stackArr)
	s.Pop()
	fmt.Println("Elements inside stack: ", s.stackArr)
	if s.Peek() != nil {
		fmt.Println(*s.Peek())
	}
	s.Push(8)
	fmt.Println(s.stackArr)
	s.Pop()
	s.Pop()
	s.Pop()
	if s.Peek() != nil {
		fmt.Println("Print no error")
	} else {

	}
}
