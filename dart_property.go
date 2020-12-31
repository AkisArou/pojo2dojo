package main

import "fmt"

const (
	DPRIVATE   = "_"
	DPUBLIC    = ""
	DPROTECTED = "@protected"
)

type DartProperty struct {
	Property
}

func NewDartProperty(pg ClassPropertyGetter) *DartProperty {
	property := &DartProperty{}
	property.SetName(pg.GetName())
	property.SetPropType(pg.GetPropType())
	property.SetAccessor(pg.GetAccessor())
	property.SetDefaultValue(pg.GetDefaultValue())

	return property
}

// ClassPropertySetter
func (p *DartProperty) SetAccessor(acc string) {
	switch acc {
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

//ClassPropertyResultProvider
func (p *DartProperty) GetPropertyString(isForConstructor bool) string {
	var propertyString string

	if p.Accessor == DPROTECTED && !isForConstructor {
		propertyString += p.Accessor + "\n\t"
	}

	if p.Accessor == DPRIVATE {
		propertyString += fmt.Sprintf("%s %s%s", p.PropType, DPRIVATE, p.Name)
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
