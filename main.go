package main

import (
	"fmt"
)

func main() {
	className, classProperties := parseJavaClass(javaClassExample)
	dartClass := makeDartClass(className, classProperties)
	fmt.Println(dartClass)
}
