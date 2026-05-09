package main

import "fmt"

var Array2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func Even(arr []int, evench chan int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			evench <- i
		}
	}
	close(evench)
}

func Odd(arr []int, oddch chan int) {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			oddch <- i
		}
	}
	close(oddch)
}
func Mix() {
	evench := make(chan int)
	oddch := make(chan int)

	go Odd(Array2,evench)
	go Even(Array2,oddch)
	fmt.Println("even numbers")
	for n := range oddch {
		fmt.Print(n,n*n)
	}
	fmt.Println("odd numbers")
	for n:=range evench{
		fmt.Print(n,n*n)
	}

}