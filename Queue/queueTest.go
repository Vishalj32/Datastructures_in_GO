package Queue

import "fmt"

func QueueRunner() {
	queue := NewQueue(10)
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)

	queue.Print()
	fmt.Println(queue.Size())
	queue.Remove()

	queue.Print()
	fmt.Println(queue.Size())

	fmt.Println(queue.Peek())
}
