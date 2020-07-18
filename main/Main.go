package main

import (
	"DataStructures/stack"
	"fmt"
)

func main()  {
	s := stack.CreateStack(make([]int, 0, 0))
	fmt.Println(s.Pop())
	s.Push(1)
	fmt.Println(s.Pop())
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
