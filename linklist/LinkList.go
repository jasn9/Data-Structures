package linklist

import "fmt"

type LinkList interface {
	GetNext() *LinkList
	SetNext(a *LinkList)
	CreateNode(a interface{}) LinkList
}

func AddAtBeginning(head LinkList, a interface{}) {
	node := head.CreateNode(a)
	node.SetNext(head.GetNext())
	head.SetNext(&node)
}

func AddAtEnd(head LinkList, a interface{}) {
	node := head.CreateNode(a)
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