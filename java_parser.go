package main

import (
	"regexp"
	"strings"
)

/* Known bugs:
1) breaks if class keyword contained in comments before class definition
2) breaks if method declaration starting bracket, is placed like C# style
3) comments are half missing
*/

func parseJavaClass(class string) (string, []string) {
	reClassName := regexp.MustCompile(`class\s(\w+)`)
	className := strings.TrimSpace(strings.Split(reClassName.FindString(class), "class")[1])

	reClassBody := regexp.MustCompile(`{[\s\S]*}`)
	classBodyStr := reClassBody.FindString(class)

	classBodyLines := strings.Split(strings.TrimSpace(classBodyStr[1:len(classBodyStr)-1]), "\n")

	var classProperties []string

	blockCount := 0

	for _, line := range classBodyLines {
		if blockCount >= 1 && strings.Contains(line, "}") {
			blockCount -= 1
			continue
		} else if blockCount >= 1 || line == "" || line == " " {
			continue
		} else if strings.Contains(line, "{") {
			blockCount += 1
			continue
		} else if strings.Contains(line, " static ") {
			continue
		} else {
			classProperties = append(classProperties, line)
		}
	}

	return className, classProperties
}

func parseJavaProperty(javaPropUnparsed string) (*JavaProperty, error) {
	javaProp := &JavaProperty{}

	parts := strings.Split(strings.TrimSpace(javaPropUnparsed), " ")

	for idx, part := range parts {
		parts[idx] = strings.Replace(strings.TrimSpace(part), ";", "", 1)
	}

	if JavaAccessors(parts[0]) != JPROTECTED &&
		JavaAccessors(parts[0]) != JPUBLIC &&
		JavaAccessors(parts[0]) != JPRIVATE {
		parts = append([]string{string(JPUBLIC)}, parts...)
	}

	javaProp.Name = parts[2]
	javaProp.PropType = parts[1]
	javaProp.Accessor = JavaAccessors(parts[0])

	hasDefaultValue := strings.Index(javaPropUnparsed, "=") > -1

	if hasDefaultValue {
		javaProp.DefaultVal = parts[len(parts)-1]
	} else {
		javaProp.DefaultVal = ""
	}

	return javaProp, nil
}
