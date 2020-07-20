package example

import "DataStructures/linklist"

func ListExample()  {
	head := linklist.DefaultLinkList{}
	linklist.AddAtBeginning(&head, &linklist.DefaultLinkList{Value: 0})
	linklist.AddAtBeginning(&head, &linklist.DefaultLinkList{Value: 2})
	linklist.AddAtEnd(&head, &linklist.DefaultLinkList{Value: 1})
	linklist.Display(&head)
}
