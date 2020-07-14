package main

import "fmt"

func main() {
	fmt.Println(Hello(""))
}

func Hello(name string) string {
	return "Hello, " + name
}
