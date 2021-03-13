package main

import "fmt"

type Hello struct {
	name string
}

func (h Hello) hello() {
	fmt.Println("Hello,", h.name, "!!!")
}

/**
* main function
**/
func main() {
	fmt.Println("Run main")
	a := &Hello{name: "Rawin"}
	a.hello()
}
