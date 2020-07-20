package linklist

import "fmt"

type LinkList interface {
	GetNext() *LinkList
	SetNext(a *LinkList)
}

func AddAtBeginning(head LinkList, node LinkList) {
	node.SetNext(head.GetNext())
	head.SetNext(&node)
}

func AddAtEnd(head LinkList, node LinkList)  {
	cur := &head
	for (*cur).GetNext()!=nil{
		cur = (*cur).GetNext()
	}
	(*cur).SetNext(&node)
}

func Display(head LinkList)  {
	cur := &head
	cur = (*cur).GetNext()
	for cur!=nil {
		fmt.Println(*cur)
		cur = (*cur).GetNext()
	}
}