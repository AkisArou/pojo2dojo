package main

func main() {
	mockDartFileName := "mockDart.dart"
	mockJavaFileName := "mockJava.java"

	mockJavaFile := readJavaFile(makePathWithFolderFile(baseFolder, mockJavaFileName))
	className, classProperties := parseJavaClass(mockJavaFile)
	dartClass := makeDartClass(className, classProperties)
	generateDartFile(dartClass, makePathWithFolderFile(baseFolder, mockDartFileName))
}
