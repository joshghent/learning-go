package main

import (
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	go fizzBuzz() // run as a goroutine, we need to sleep the thread or otherwise the app will quit before the method has been called

	println("this comes first")

	time.Sleep(100 * time.Millisecond)
}

func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			go println("FizzBuzz")
		} else if i%3 == 0 {
			go println("Fizz")
		} else if i%5 == 0 {
			go println("Buzz")
		}
	}
}
