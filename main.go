package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	pf      = fmt.Printf
	pl      = fmt.Println
	balance int
)

func main() {
	pl("Mutex Example")

	done := make(chan bool)
	go withdraw(700, done)
	go deposit(500, done)

	<-done
	<-done

	pl("final balance", balance)
}

func init() {
	balance = 1000
}

func deposit(value int, done chan bool) {
	mutex.Lock()
	balance += value
	mutex.Unlock()
	done <- true
	pl("deposit", balance)
}

func withdraw(value int, done chan bool) {
	mutex.Lock()
	balance -= value
	mutex.Unlock()
	done <- true
	pl("witdraw", balance)
}
