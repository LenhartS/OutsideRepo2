package Repo2

import Repo4 "github.com/LenhartS/OutsideRepo4"

var s4 string

func SaySomething2() string {
	println("String from GolangOutsieRepo2")
	s4 := Repo4.SaySomething4()
	println("You got a present from Repo4")
	println(s4, "  <-- this here came from Repo4 to Repo2")
	return s4
}
