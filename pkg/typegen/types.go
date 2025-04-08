package typegen

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

type NamingConvention string

const (
	NamingConventionCamelCase  NamingConvention = "camelCase"
	NamingConventionPascalCase NamingConvention = "PascalCase"
	NamingConventionSnakeCase  NamingConvention = "snake_case"
)

type TypeDefinition struct {
	Name       string
	Properties []Property
	IsExported bool
}

type Property struct {
	Name     string
	Type     string
	Required bool
	Tags     map[string]string
}
