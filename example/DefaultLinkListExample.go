package example

import "DataStructures/linklist"

func ListExample()  {
	head := linklist.DefaultLinkList{}
	linklist.AddAtBeginning(&head, 0)
	linklist.AddAtBeginning(&head,  2)
	linklist.AddAtEnd(&head, 1)
	linklist.Display(&head)
}
