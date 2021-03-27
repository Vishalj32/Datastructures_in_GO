package Stack

import "fmt"

func StackRunner() {
	stack := NewStack(10)
	stack.Push(1)
	stack.Push(2)

	stack.Push(3)

	stack.Print()
	stack.Pop()
	stack.Print()
	fmt.Println("Size := ", stack.Size())
}
