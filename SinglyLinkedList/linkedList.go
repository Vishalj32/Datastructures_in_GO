package singlelinkedlist

import (
	"Datastructures_in_GO/SinglyLinkedList/src"
	"fmt"
)

//RunLinkedList executes LinkedList function
func RunLinkedList() {
	ll := src.NewLinkedList()

	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)

	ll.Print()
	fmt.Println("Size : ", ll.Length())

	ll.InsertAtPosition(4, 2)
	ll.Print()

	fmt.Println(ll.GetValue(3))
	err := ll.Delete(2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ll.Print()
	}

	err = ll.DeleteAtPosition(2)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		ll.Print()
	}
}
