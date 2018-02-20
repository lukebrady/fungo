package main

import (
	"fmt"
)

// Stack is the stack object
type Stack struct {
	stack []int
}

// Push pushes an item into the stack
func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
	fmt.Printf("%v\n", s.stack)
}

// Pop pops a value from the stack
func (s *Stack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
}

// PrintStack prints stack to STDOUT
func (s *Stack) PrintStack() {
	t := s.stack[0:]
	fmt.Printf("%v\n", t)
}

func main() {
	stack := new(Stack)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Pop()
	stack.Pop()
	stack.PrintStack()
	hello := make(chan string)
	go func(hello chan string) {
		yo := "Hello how are you?"
		hello <- yo
	}(hello)

	fmt.Println(<-hello)
}
