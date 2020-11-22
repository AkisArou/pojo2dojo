package main

import "fmt"

type DartAccessors string

const (
	DPRIVATE   DartAccessors = "_"
	DPUBLIC                  = ""
	DPROTECTED               = "@protected"
)

type DartProperty struct {
	Name       string
	PropType   string
	Accessor   DartAccessors
	DefaultVal string
}

func (p *DartProperty) SetAccessor(javaAccessor JavaAccessors) {
	switch javaAccessor {
	case JPUBLIC:
		p.Accessor = DPUBLIC
	case JPRIVATE:
		p.Accessor = DPRIVATE
	case JPROTECTED:
		p.Accessor = DPROTECTED
	}
}

func (p *DartProperty) SetDefaultValue(defaultVal string) {
	p.DefaultVal = defaultVal
}

func (p *DartProperty) SetName(name string) {
	p.Name = name
}

func (p *DartProperty) SetPropType(propType string) {
	p.PropType = propType
}

func (p *DartProperty) GetDartPropertyString(isForConstructor bool) string {
	var propertyString string

	if p.Accessor == DPROTECTED && !isForConstructor {
		propertyString += string(p.Accessor) + "\n\t"
	}

	if p.Accessor == DPRIVATE {
		propertyString += fmt.Sprintf("%s %s%s", p.PropType, string(DPRIVATE), p.Name)
	} else {
		propertyString += fmt.Sprintf("%s %s", p.PropType, p.Name)
	}

	if p.DefaultVal != "" {
		propertyString += fmt.Sprintf(" = %s", p.DefaultVal)
	}

	if !isForConstructor {
		propertyString += ";"
	}

	return propertyString
}

func makeDartProperty(javaParts *[4]string) *DartProperty {
	property := DartProperty{}
	property.SetName(javaParts[0])
	property.SetPropType(javaParts[1])
	property.SetAccessor(JavaAccessors(javaParts[2]))
	property.SetDefaultValue(javaParts[3])

	return &property
}
