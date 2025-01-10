package main

import "fmt"
import github.com/LenhartS/OutsideRepo3/main.go

func main() {
	fmt.Println("This file is in GoLangOutsideRepo2 ")
	fmt.Println(SaySomething2())

}

func SaySomething2() string {
	s1 := "String from GolangOutsieRepo2"
	return s1
}
