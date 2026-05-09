package main

import "fmt"

func Queue() {
	q := []int{}

	// enqueue
	q = append(q, 10)
	q = append(q, 20)
	q = append(q, 30)
	q = append(q, 40)
	fmt.Println(q)

	// dequeue
	q=q[1:]
	fmt.Println(q)

	fmt.Println(q[0])
}