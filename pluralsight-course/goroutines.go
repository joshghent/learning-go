package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	ch := make(chan string)
	doneCh := make(chan bool) // channel for when it is done
	go abc(ch)                // run as a goroutine, we need to sleep the thread or otherwise the app will quit before the method has been called
	go printer(ch, doneCh)

	// Receive the message on the done channel
	// Main will stay alive until this completes
	<-doneCh
}

func abc(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		// Pass the string to the channel
		ch <- string(l)
	}

	// Close the channel so it will be no longer available
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	// Wait for messages to come in on the channel
	for l := range ch {
		println(l)
	}

	// Pass true message on the done channel
	doneCh <- true
}
