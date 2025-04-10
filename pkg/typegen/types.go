package typegen

type NamingConvention string

const (
	NamingConventionCamelCase NamingConvention = "camelCase"
	NamingConventionSnakeCase NamingConvention = "snake_case"
)

type Generator struct {
	InputDir  string
	OutputDir string
	Options   *Options
}

type Options struct {
	UseInterface     bool
	UseReadonly      bool
	OutputFileName   string
	NamingConvention NamingConvention
}

type TypeDefinition struct {
	Name       string
	Package    string
	Properties []Property
	IsExported bool
}

type Property struct {
	Name     string
	Type     string
	Required bool
	Tags     map[string]string
}
