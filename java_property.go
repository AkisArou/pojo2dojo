package main

const (
	JPRIVATE   = "private"
	JPUBLIC    = "public"
	JPROTECTED = "protected"
)

type JavaProperty struct {
	Property
}

// ClassPropertySetter
func (jp *JavaProperty) SetAccessor(acc string) {
	jp.Accessor = acc
}

func (jp *JavaProperty) SetDefaultValue(defaultVal string) {
	jp.DefaultVal = defaultVal
}

func (jp *JavaProperty) SetName(name string) {
	jp.Name = name
}

func (jp *JavaProperty) SetPropType(propType string) {
	jp.PropType = propType
}

//ClassPropertyGetter
func (jp *JavaProperty) GetAccessor() string {
	return jp.Accessor
}

func (jp *JavaProperty) GetDefaultValue() string {
	return jp.DefaultVal
}

func (jp *JavaProperty) GetName() string {
	return jp.Name
}

func (jp *JavaProperty) GetPropType() string {
	return jp.PropType
}

func (jp *JavaProperty) GetPropertyString(isForConstructor bool) string {
	return "To be implemented"
}
