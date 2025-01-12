package Repo2

func SaySomething2() string {
	s1 := "String from GolangOutsieRepo2"
	return s1
}

func SaySomethingfromRepo3() string {
	s2 := Repo3Pkg.SaySomething3()
	return s2
}
