package stack

import "fmt"

type Stack struct {
	stack []int
}

type Error string

func (err Error) Error() string{
	return fmt.Sprintf("StackError: %s",string(err))
}

func (s *Stack) Pop() (int, error) {
	if len(s.stack)==0{
		return -1, Error("stack Underflow")
	}
	ele := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return ele, nil
}

func (s *Stack) Push(a int) error {
	s.stack = append(s.stack, a)
	return nil
}

func CreateStack(a []int) Stack {
	return Stack{a}
}

