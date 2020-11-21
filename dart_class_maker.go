package main

import (
	"fmt"
	"os"
	"strings"
)

const class = "class"
const requiredAnnotation = "@required"

func makeDartClass(className string, javaClassProperties []string) string {
	var dartProperties string
	var dartConstructorProperties string

	for _, prop := range javaClassProperties {
		if prop == "" || prop == " " {
			continue
		}

		javaParts, err := parseJavaProperty(prop)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		dartProp := makeDartProperty(javaParts)
		dartPropertyString := dartProp.getDartPropertyString(false)
		dartProperties += "\t" + dartPropertyString + "\n"
		dartConstructorProperties += "\t\t" + requiredAnnotation + " " + dartProp.getDartPropertyString(true) + ",\n"
	}

	dartConstructorProperties = strings.TrimSpace(dartConstructorProperties)

	dartClass := fmt.Sprintf(`
%s %s {
%s
	%s({
		%s
	})
}
`, class, className, dartProperties, className, dartConstructorProperties)
	////output := "class " + className + " {\n" + dartProperties + "\n\t" + className + "\n}"
	//return output
	return dartClass
}
