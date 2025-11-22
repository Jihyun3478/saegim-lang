package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	environment := NewEnvironment()
	environment.outer = outer
	return environment
}

func (environment *Environment) Set(name string, value Object) Object {
	environment.store[name] = value
	return value
}

func (environment *Environment) Get(name string) (Object, bool) {
	object, ok := environment.store[name]
	if !ok && environment.outer != nil {
		object, ok = environment.outer.Get(name)
	}
	return object, ok
}
