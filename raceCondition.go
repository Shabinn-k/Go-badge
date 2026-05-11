package main

import (
	"fmt"
	"sync"
)

var Bank int
var mu sync.Mutex

func Deposit(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	Bank+=10
	mu.Unlock()
}

func Race() {
	var wg sync.WaitGroup
	for i:=0;i<10000;i++{
		wg.Add(1)
		go Deposit(&wg)
	}
	wg.Wait()
	fmt.Println(Bank)
}