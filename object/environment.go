package object

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

func (environment *Environment) Set(name string, value Object) Object {
	environment.store[name] = value
	return value
}

func (environment *Environment) Get(name string) (Object, bool) {
	object, ok := environment.store[name]
	return object, ok
}
