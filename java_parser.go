package main

import (
	"fmt"
	"regexp"
	"strings"
)

type JavaParser struct{}

func (jp *JavaParser) Parse(class string) *ParsedResult {
	classBodyLines := strings.Split(class, "\n")

	var classProperties []string
	var classNameContainingLine string

	blockCount := 0

	for _, line := range classBodyLines {
		if jp.isComment(line) {
			continue
		} else if jp.isMethodAnnotation(line) {
			continue
		} else if jp.isEmptyLine(line, blockCount) {
			continue
		} else if jp.isClassNameContainingLine(line) {
			classNameContainingLine = line
		} else if jp.isStartingBlock(line) {
			blockCount += 1
		} else if jp.isEndingBlock(line, blockCount) {
			blockCount -= 1
		} else if jp.isMethod(line) {
			continue
		} else {
			classProperties = append(classProperties, line)
		}
	}

	pr := &ParsedResult{}
	pr.className = jp.setClassName(classNameContainingLine)
	pr.classProperties = jp.getClassProperties(classProperties)

	return pr
}

func (jp *JavaParser) isComment(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "//")
}

func (jp *JavaParser) isMethodAnnotation(line string) bool {
	return strings.HasPrefix(strings.TrimSpace(line), "@")
}

func (jp *JavaParser) isClassNameContainingLine(line string) bool {
	reClassName := regexp.MustCompile(`class\s(\w+)`)
	return reClassName.MatchString(line)
}

func (jp *JavaParser) isStartingBlock(line string) bool {
	return strings.HasSuffix(line, "{") || strings.HasPrefix(line, "{")
}

func (jp *JavaParser) isEndingBlock(line string, blockCount int) bool {
	return blockCount >= 1 && (strings.HasSuffix(line, "}") || strings.HasPrefix(line, "}"))
}

func (jp *JavaParser) isEmptyLine(line string, blockCount int) bool {
	return blockCount >= 1 || line == "" || line == " "
}

func (jp *JavaParser) isMethod(line string) bool {
	reMethod := regexp.MustCompile(`\(([^)]+)\)`)
	return reMethod.MatchString(line)
}

func (jp *JavaParser) setClassName(classNameContainingLine string) string {
	s := strings.Split(classNameContainingLine, classKeyword)[1]
	return strings.Split(s, " ")[1]
}

func (jp *JavaParser) getClassProperties(rawProps []string) []ClassPropertySGP {
	var cp []ClassPropertySGP

	for _, prop := range rawProps {
		if strings.Contains(prop, ",") {
			var pProps []string
			var defaultValue = ""

			multiProps := strings.Split(prop, ",")
			multiPropsLen := len(multiProps)
			indexProp := strings.Split(multiProps[0], " ")
			lastProp := multiProps[multiPropsLen-1]
			defaultValueIdx := strings.Index(lastProp, "=")

			if defaultValueIdx > -1 {
				defaultValue = lastProp[strings.Index(lastProp, "="):]
				multiProps[multiPropsLen-1] = strings.Replace(lastProp, defaultValue, "", 1)
			}

			propAccType := strings.Join(indexProp[:len(indexProp)-1], " ")
			pProps = append(pProps, fmt.Sprintf("%s %s", multiProps[0], defaultValue))

			for i := 1; i < multiPropsLen; i++ {
				pProps = append(pProps, fmt.Sprintf("%s %s %s", propAccType, strings.TrimSpace(multiProps[i]), defaultValue))
			}

			for _, s := range pProps {
				cp = append(cp, jp.parseRawProperty(s))
			}

		} else {
			cp = append(cp, jp.parseRawProperty(prop))
		}
	}

	return cp
}

func (jp *JavaParser) parseRawProperty(propertyRaw string) ClassPropertySGP {
	javaPropUnparsed := strings.Replace(strings.TrimSpace(propertyRaw), ";", "", 1)
	javaProp := &JavaProperty{}

	parts := strings.Split(javaPropUnparsed, " ")

	if parts[0] != JPROTECTED &&
		parts[0] != JPUBLIC &&
		parts[0] != JPRIVATE {
		parts = append([]string{JPUBLIC}, parts...)
	}

	javaProp.SetName(parts[2])
	javaProp.SetPropType(parts[1])
	javaProp.SetAccessor(parts[0])

	hasDefaultValue := strings.Index(javaPropUnparsed, "=") > -1

	if hasDefaultValue {
		javaProp.SetDefaultValue(parts[len(parts)-1])
	} else {
		javaProp.SetDefaultValue("")
	}

	return javaProp
}
