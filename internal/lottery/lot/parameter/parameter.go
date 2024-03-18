package parameter

type Option struct {
	patterns []string
}

type option func(*Option)

type Parameter struct {
	path string
	Option
}

func New(path string, option ...option) *Parameter {
	parameter := &Parameter{
		path: path,
		Option: Option{
			patterns: []string{"image/"},
		},
	}

	for _, value := range option {
		value(&parameter.Option)
	}

	return parameter
}

func (p Parameter) Path() string {
	return p.path
}

func (p Parameter) Patterns() []string {
	return p.patterns
}
