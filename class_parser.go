package main

type ClassParser interface {
	isComment(line string) bool
	isClassNameContainingLine(line string) bool
	isStartingBlock(line string) bool
	isEndingBlock(line string, blockCount int) bool
	isEmptyLine(line string, blockCount int) bool
	isMethod(line string) bool
	isMethodAnnotation(line string) bool

	setClassName(classNameContainingLine string) string
	parseRawProperty(propertyRaw string) *Property

	Parse() *ParsedResult
}
