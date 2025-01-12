package Repo2

import (
	"github.com/LenhartS/OutsideRepo3/"
)

func SaySomething2() string {
	s1 := "String from GolangOutsieRepo2"
	return s1
}

func SaySomethingfromRepo3() string {
	s2 := OutsideRepo3.SaySomething3()
	return s2
}
