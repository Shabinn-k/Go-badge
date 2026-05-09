package main

import "fmt"

func Stack() {
	stk := []int{}
	// push
	stk = append(stk, 1)
	stk = append(stk, 2)
	stk = append(stk, 3)
	stk = append(stk, 4)
	fmt.Println(stk)

	// peek
	fmt.Println(stk[len(stk)-1])
	
	// pop
	stk=stk[:len(stk)-1]
	fmt.Println(stk)
}
