package main

import "fmt"

func Queue() {
	q := []int{}

	// enqueue
	q = append(q, 100)
	q = append(q, 200)
	q = append(q, 300)
	q = append(q, 400)
	fmt.Println(q)

	// dequeue
	q=q[1:]
	fmt.Println(q)

	fmt.Println(q[0])
}

// func main(){ 
//     Queue()
//     fmt.Println("OG :",q)
    
//     Enq(4)
//     fmt.Println("After ENQ :",q)
    
//     Deq()
//     fmt.Println("After DEQ :",q)
    
//     Peek()
        
// }

// var q=[]int{}
// func Queue(){
//     q = append(q, 1)
// 	q = append(q, 2)
// 	q = append(q, 3)
// }

// func Enq(num int)[]int{
//     q=append(q,num) 
//     return q
// }

// func Peek(){
//     get:=q[0]
//     fmt.Println("Peek element :",get)
// }

// func Deq(){
//     q=q[1:]
// }

