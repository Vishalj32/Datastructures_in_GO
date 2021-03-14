package Queue

import "fmt"

func QueueRunner() {
	queue := NewQueue(10)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	queue.Print()
	fmt.Println("Size := ", queue.Size())
	queue.Remove()

	queue.Print()
	fmt.Println("Size := ", queue.Size())

	fmt.Println("Peek := ", queue.Peek())
	fmt.Println("Size := ", queue.Size())
}
