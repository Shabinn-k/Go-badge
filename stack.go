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


// func main(){ 
//     Stack()
    
//     Push(6)
//     fmt.Println(stk)

//     Pop()
//     fmt.Println(stk)
    
//     Peek()
// }

// var stk=[]int{}

// func Stack(){
//     // push
//     stk=append(stk,1)
//     stk=append(stk,2)
//     stk=append(stk,3)
//     stk=append(stk,4)
//     stk=append(stk,5)
    
// }

// func Pop(){
//     stk=stk[:len(stk)-1]
// }
// func Peek(){
//     get:=stk[len(stk)-1]
//     fmt.Println(get)
// }
// func Push(num int)[]int{
//     stk=append(stk,num)
//     return stk
// }