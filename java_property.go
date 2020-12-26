package main

type JavaAccessors string

const (
	JPRIVATE   JavaAccessors = "private"
	JPUBLIC                  = "public"
	JPROTECTED               = "protected"
)

type JavaProperty struct {
	Name       string
	PropType   string
	Accessor   JavaAccessors
	DefaultVal string
}
