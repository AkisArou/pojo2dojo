package main

func main() {
	mockDartFileName := "mockDart.dart"
	mockJavaFileName := "mockJava.java"

	javaParser := JavaParser{}
	dartBuilder := DartClassBuilder{}

	mockJavaFile := readJavaFile(makePathWithFolderFile(baseFolder, mockJavaFileName))

	result := javaParser.Parse(mockJavaFile)
	dc := dartBuilder.Build(result)
	generateDartFile(dc, makePathWithFolderFile(baseFolder, mockDartFileName))
}
