package main

import (
	"fmt"
	"regexp"
	"strings"
)

type JavaAccessors string

const (
	JPRIVATE   JavaAccessors = "private"
	JPUBLIC                  = "public"
	JPROTECTED               = "protected"
)

func parseJavaClass(class string) (string, []string) {
	reClassName := regexp.MustCompile(`class\s(\w+)`)
	className := strings.TrimSpace(strings.Split(reClassName.FindString(class), "class")[1])

	reClassProperties := regexp.MustCompile(`{[\s\S]*}`)
	classPropertiesStr := reClassProperties.FindString(class)
	classProperties := strings.Split(strings.TrimSpace(classPropertiesStr[1:len(classPropertiesStr)-1]), ";")

	return className, classProperties
}

func parseJavaProperty(javaProp string) (*[4]string, error) {
	if javaProp == "" || javaProp == " " {
		return &[4]string{}, fmt.Errorf("not parsable string")
	}

	parts := strings.Split(strings.TrimSpace(javaProp), " ")

	for idx, part := range parts {
		parts[idx] = strings.TrimSpace(part)
	}

	hasDefaultValue := strings.Index(javaProp, "=") > -1

	var defaultVal string

	if hasDefaultValue {
		defaultVal = parts[len(parts)-1]
	} else {
		defaultVal = ""
	}

	return &[4]string{parts[2], parts[1], parts[0], defaultVal}, nil
}
