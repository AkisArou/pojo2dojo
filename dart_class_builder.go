package main

import (
	"fmt"
	"strings"
)

const classKeyword = "class"
const requiredAnnotation = "@required"

type DartClassBuilder struct{}

func (d *DartClassBuilder) Build(result *ParsedResult) string {
	var dartProperties string
	var dartConstructorProperties string

	for _, prop := range result.classProperties {
		dp := NewDartProperty(prop)
		dps := dp.GetPropertyString(false)
		dartProperties += "\t" + dps + "\n"
		dartConstructorProperties += "\t\t" + requiredAnnotation + " " + dp.GetPropertyString(true) + ",\n"
	}
	dartConstructorProperties = strings.TrimSpace(dartConstructorProperties)

	dartClass := fmt.Sprintf(`
%s %s {
%s
	%s({
		%s
	})
}
`, classKeyword, result.className, dartProperties, result.className, dartConstructorProperties)
	return dartClass

}
