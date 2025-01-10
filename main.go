package main

import "fmt"

func main() {
	fmt.Println("This file is in GoLangOutsideRepo2 ")
	fmt.Println(SaySomething())
}

func SaySomething() string {
	s1 := "String from GolangOutsieRepo2"
	return s1
}
