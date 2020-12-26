package main

type JavaAccessors string

const (
	JPRIVATE   JavaAccessors = "private"
	JPUBLIC    JavaAccessors = "public"
	JPROTECTED JavaAccessors = "protected"
)

type JavaProperty struct {
	Name       string
	PropType   string
	Accessor   JavaAccessors
	DefaultVal string
}
