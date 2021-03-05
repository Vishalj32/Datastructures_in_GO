package src

import (
	"errors"
	"fmt"
)

//Node :- stores the internal linkedlist data
//
//value :- node value
//
//next :- pointer to the next element in the LinkedList
type Node struct {
	value int
	next  *Node
}

//LinkedList datastructure
//
//Stores the length of the LinkedList at variable 'len'
//
//Stores the pointer to the head node
type LinkedList struct {
	head *Node
	len  int
}

//NewLinkedList := creates a new LinkedList and return's a pointer for the same
func NewLinkedList() *LinkedList {
	linkedList := new(LinkedList)
	return linkedList
}

//GetValue := returns the value at the provided position
func (linkedList *LinkedList) GetValue(position int) int {
	node := linkedList.Get(position)
	return node.value
}

//Get := retrieves the entry at the provided position
func (linkedList *LinkedList) Get(poistion int) *Node {

	head := linkedList.head

	if poistion < 0 || poistion > (linkedList.len-1) {
		fmt.Println("Invalid position provided")
	}

	for i := 0; i < poistion; i++ {
		head = head.next
	}
	return head
}

//Insert :- insert's the new node at the end of the LinkedList
func (linkedList *LinkedList) Insert(value int) {
	node := Node{}
	node.value = value

	if linkedList.len == 0 {
		linkedList.head = &node
		linkedList.len++
		return
	}

	head := linkedList.head
	for i := 0; i < linkedList.len; i++ {
		if head.next == nil {
			head.next = &node
			linkedList.len++
			return
		}
		head = head.next
	}
}

//InsertAtPosition := Insert's the given value at the position
func (linkedList *LinkedList) InsertAtPosition(value int, position int) {
	node := Node{}
	node.value = value

	if position < 0 {
		fmt.Println("Invalid position")
		return
	}

	if position == 0 {
		linkedList.head = &node
		linkedList.len++
		return
	}

	if position > linkedList.len {
		fmt.Println("Position out of bounds")
		return
	}

	temp := linkedList.Get(position)
	node.next = temp
	prevNode := linkedList.Get(position - 1)
	prevNode.next = &node
	linkedList.len++

}

//DeleteAtPosition := removes the linkedlist node at the given position
func (linkedList *LinkedList) DeleteAtPosition(position int) error {
	if position < 0 || position > (linkedList.len-1) {
		fmt.Println("Invalid position")
		return errors.New("Invalid position")
	}

	if linkedList.len == 0 {
		fmt.Println("Empty LinkedList")
		return errors.New("Empty LinkedList, Delete failed")
	}

	previousNode := linkedList.Get(position - 1)

	if previousNode == nil {
		fmt.Println("Empty LinkedList")
		return errors.New("Empty LinkedList, Delete failed")
	}
	previousNode.next = linkedList.Get(position).next
	linkedList.len--
	return nil
}

//Delete := removes the provided value from the LinkedList
func (linkedList *LinkedList) Delete(value int) error {
	head := linkedList.head

	if linkedList.len == 0 {
		fmt.Println("Empty LinkedList")
		return errors.New("Empty LinkedList, Delete failed")
	}

	for i := 0; i < linkedList.len; i++ {
		if head.value == value {
			if i > 0 {
				prevNode := linkedList.Get(i - 1)
				prevNode.next = linkedList.Get(i).next
			} else {
				linkedList.head = head.next
			}
			linkedList.len--
			return nil
		}
		head = head.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}

//Print :- prints the nodes in LinkedList
func (linkedList *LinkedList) Print() {
	if linkedList.len == 0 {
		fmt.Println("Empty LinkedList!!")
	}

	head := linkedList.head
	for i := 0; i < linkedList.len; i++ {
		if head.next == nil {
			fmt.Print(head.value)
		} else {
			fmt.Print(head.value, "->")
		}
		head = head.next
	}
	fmt.Println()
}

//Length :-  returns the length of the linkedList
func (linkedList *LinkedList) Length() int {
	return linkedList.len
}
