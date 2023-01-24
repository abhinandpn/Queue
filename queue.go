package main

import (
	"fmt"
	"os"
	"strconv"
)

type node struct {
	number int
	next   *node
}
type queue struct {
	len   int
	front *node
	rear  *node
}

func (head *queue) Enqueue(num int) {
	var temp = &node{number: num}
	// temp.number = num
	// temp.next = nil
	if head.front == nil && head.rear == nil {
		head.front = temp
		head.rear = temp
	} else {
		head.rear.next = temp
		head.rear = temp
	}
	head.len++
}
func (head *queue) Dequeue() {
	if head.rear == head.front {
		head.rear = nil
		head.front = nil
	} else {
		head.front = head.front.next
	}
	head.len--
}
func (head *queue) Front() {
	fmt.Println("Front node is ", head.front.number)
}
func (head *queue) Size() {
	fmt.Println("Size of queue :", head.len)
}
func (head *queue) Display() {
	var p *node = head.front
	for p != nil {
		fmt.Printf("%d <- ", p.number)
		p = p.next
	}
}
func main() {
	var myQueue = &queue{}
	var choice string
	for {
		fmt.Println("\n  Enter your choice :")
		fmt.Println("1. Enqueue")
		fmt.Println("2. Dequeue")
		fmt.Println("3. Front ")
		fmt.Println("4. Size")
		fmt.Println("5. Traverse/Display")
		fmt.Println("6. Exit")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			var data string
			fmt.Println("Enter the Number for node : ")
			fmt.Scanln(&data)
			num, _ := strconv.Atoi(data)
			myQueue.Enqueue(num)
		case "2":
			myQueue.Dequeue()
		case "3":
			myQueue.Front()
		case "4":
			myQueue.Size()
		case "5":
			myQueue.Display()
		default:
			os.Exit(0)
		}
	}
}
