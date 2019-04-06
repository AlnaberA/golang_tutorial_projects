package main

// https://golang.org/pkg/

import "golang_tutorial_projects/util"
import "golang_tutorial_projects/util/helper"
import "golang_tutorial_projects/mathPackage"

//Aliases can be used to simplify import usuage
//import help "../util/helper"

//Wont work because util "redeclared as imported package name previous declaration at line 5"
//import util "../util/helper"
func main() {
	println(util.GetPackageName())
	helper.PrintHelloWorld()
	var num = mathPackage.Add(1, 1)
	println(num)
}
