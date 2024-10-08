package main

import (
	"fmt"
	"sync"
)

type Stack struct {
	stackArr []int
	sMux     sync.Mutex
}

func (s *Stack) Push(ele int) {
	s.sMux.Lock()
	s.stackArr = append(s.stackArr, ele)
	s.sMux.Unlock()
}

func (s *Stack) Pop() {
	s.sMux.Lock()
	defer s.sMux.Unlock()
	if len(s.stackArr) <= 0 {
		fmt.Println("Underflow")
		return
	}
	tempArr := s.stackArr[:len(s.stackArr)-1]
	s.stackArr = tempArr
}

func (s *Stack) Peek() int {
	s.sMux.Lock()
	defer s.sMux.Unlock()
	return s.stackArr[len(s.stackArr)-1]
}

func main() {
	// example 1 to test out functionality
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
	// example 2 to test out functionality
	// s := Stack{}
	// s.Pop()
}
