package main

// https://golang.org/pkg/

import "../util"
import "../util/helper"
import "packaging/math_package"

//Aliases can be used to simplify import usuage
//import help "../util/helper"

//Wont work because util "redeclared as imported package name previous declaration at line 5"
//import util "../util/helper"
func main() {
	println(util.GetPackageName())
	helper.PrintHelloWorld()
	math_package.Add(1, 1)
}
