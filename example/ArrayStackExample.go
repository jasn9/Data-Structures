package example

import (
	"DataStructures/stack"
	"fmt"
)

func Example()  {
	s := stack.ArrayStack{S: make([]interface{}, 0, 0)}
	val, err := stack.Pop(&s)
	if err==nil{
		fmt.Println(val)
	} else{
		fmt.Println(err)
	}

	stack.Push(&s, 1)
	fmt.Println(s)
	val, err = stack.Pop(&s)
	if err==nil{
		fmt.Println(val)
	} else{
		fmt.Println(err)
	}

	stack.Push(&s, 2)
	fmt.Println(s)
	stack.Push(&s, "abc")
	fmt.Println(s)
}
