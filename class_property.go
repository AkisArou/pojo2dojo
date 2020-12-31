package main

type Property struct {
	Name       string
	PropType   string
	Accessor   string
	DefaultVal string
}

type ClassPropertySetter interface {
	SetAccessor(acc string)
	SetDefaultValue(defaultVal string)
	SetName(name string)
	SetPropType(propType string)
}

type ClassPropertyGetter interface {
	GetAccessor() string
	GetDefaultValue() string
	GetName() string
	GetPropType() string
}

type ClassPropertyResultProvider interface {
	//TODO isForConstructor
	GetPropertyString(isForConstructor bool) string
}

type ClassPropertySGP interface {
	ClassPropertySetter
	ClassPropertyGetter
	ClassPropertyResultProvider
}
