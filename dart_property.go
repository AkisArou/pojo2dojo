package main

import "fmt"

type DartAccessors string

const (
	DPRIVATE   DartAccessors = "_"
	DPUBLIC                  = ""
	DPROTECTED               = "@protected"
)

type DartProperty struct {
	name       string
	propType   string
	accessor   DartAccessors
	defaultVal string
}

func (p *DartProperty) setAccessor(javaAccessor JavaAccessors) {
	switch javaAccessor {
	case JPUBLIC:
		p.accessor = DPUBLIC
	case JPRIVATE:
		p.accessor = DPRIVATE
	case JPROTECTED:
		p.accessor = DPROTECTED
	}
}

func (p *DartProperty) setDefaultValue(defaultVal string) {
	p.defaultVal = defaultVal
}

func (p *DartProperty) setName(name string) {
	p.name = name
}

func (p *DartProperty) setPropType(propType string) {
	p.propType = propType
}

func (p *DartProperty) getDartPropertyString(isForConstructor bool) string {
	var propertyString string

	if p.accessor == DPROTECTED && !isForConstructor {
		propertyString += string(p.accessor) + "\n\t"
	}

	if p.accessor == DPRIVATE {
		propertyString += fmt.Sprintf("%s %s%s", p.propType, string(DPRIVATE), p.name)
	} else {
		propertyString += fmt.Sprintf("%s %s", p.propType, p.name)
	}

	if p.defaultVal != "" {
		propertyString += fmt.Sprintf(" = %s", p.defaultVal)
	}

	if !isForConstructor {
		propertyString += ";"
	}

	return propertyString
}

func makeDartProperty(javaParts *[4]string) *DartProperty {
	property := DartProperty{}
	property.setName(javaParts[0])
	property.setPropType(javaParts[1])
	property.setAccessor(JavaAccessors(javaParts[2]))
	property.setDefaultValue(javaParts[3])

	return &property
}
