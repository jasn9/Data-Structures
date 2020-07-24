package linklist

type DefaultLinkList struct{
	Value interface{}
	Next *LinkList
}

func (list *DefaultLinkList) GetNext()  *LinkList {
	return list.Next
}

func (list *DefaultLinkList) SetNext(a *LinkList)  {
	list.Next = a
}

func (list *DefaultLinkList) CreateNode(a interface{}) LinkList{
	return &DefaultLinkList{a, nil}
}