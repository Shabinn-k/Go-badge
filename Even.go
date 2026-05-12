package main

import "fmt"

func Even(evench chan int) {
	for i := 0; i <= 10; i += 2 {
		evench <- i
	}
	close(evench)
}

func Odd(oddch chan int) {
	for i := 1; i <= 10; i += 2 {
		oddch <- i
	}
	close(oddch)
}

func Mix() {
	evench := make(chan int)
	oddch := make(chan int)

	go Odd(evench)
	go Even(oddch)
	fmt.Println("even numbers")
	for n := range oddch {
		fmt.Println(n)
	}
	fmt.Println("odd numbers")
	for n := range evench {
		fmt.Println(n)
	}

}
