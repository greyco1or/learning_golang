package main

import "fmt"

type persion struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood:= []string{"kimchi", "hamburger"}
	grey := persion{name: "grey", age: 80, favFood: favFood}
	fmt.Println(grey)
}