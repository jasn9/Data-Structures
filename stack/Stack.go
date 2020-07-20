package stack

import "fmt"

type Stack interface {
	IsEmpty() bool
	RemoveLastElementAdded() interface{}
	AddElement(a interface{}) error
}

type Error string

func (err Error) Error() string{
	return fmt.Sprintf("StackError: %s",string(err))
}

func Pop(stack Stack) (interface{}, error) {
	 if stack.IsEmpty(){
		return -1, Error("stack Underflow")
	}
	return stack.RemoveLastElementAdded(), nil
}

func Push(stack Stack,a interface{}) error {
	return stack.AddElement(a)
}

