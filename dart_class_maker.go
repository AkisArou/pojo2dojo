package main

import (
	"fmt"
	"log"
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
			log.Fatal(err)
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
	return dartClass
}
