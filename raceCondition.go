package main

import (
	"fmt"
	"sync"
)

var Bank int
var mus sync.Mutex

func Deposit(wg *sync.WaitGroup) {
	defer wg.Done()
	mus.Lock()
	Bank+=10
	mus.Unlock()
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