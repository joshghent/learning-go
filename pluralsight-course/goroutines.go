package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	go abc() // run as a goroutine, we need to sleep the thread or otherwise the app will quit before the method has been called

	println("this comes first")

	time.Sleep(100 * time.Millisecond)
}

func abc() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
