package main

func main() {
	foo := myStruct{}
	foo.name = "Josh"

	println(foo.name)
}

type myStruct struct {
	name string
}
