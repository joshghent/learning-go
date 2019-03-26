package main

func main() {
	// sayInfiniteHello("hi", "whats up", "howdy")
	length, result := add(1, 2, 3, 5, 6, 7)
	println(length, result)
}

func sayHello(message *string) {
	println(*message)
	*message = "Hello"
}

func sayInfiniteHello(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}

func add(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}
