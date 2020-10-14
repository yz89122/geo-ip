package container

import "go.uber.org/dig"

// container for providing IoC
var container = dig.New()

func init() {
	container.Provide(func() *dig.Container {
		return container
	})
}

// Get the Container.
func Get() *dig.Container {
	return container
}

// Provide teaches the container how to build values of one or more types and
// expresses their dependencies.
//
// The first argument of Provide is a function that accepts zero or more
// parameters and returns one or more results. The function may optionally
// return an error to indicate that it failed to build the value. This
// function will be treated as the constructor for all the types it returns.
// This function will be called AT MOST ONCE when a type produced by it, or a
// type that consumes this function's output, is requested via Invoke. If the
// same types are requested multiple times, the previously produced value will
// be reused.
//
// In addition to accepting constructors that accept dependencies as separate
// arguments and produce results as separate return values, Provide also
// accepts constructors that specify dependencies as dig.In structs and/or
// specify results as dig.Out structs.
func Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return container.Provide(constructor, opts...)
}

// Invoke runs the given function after instantiating its dependencies.
//
// Any arguments that the function has are treated as its dependencies. The
// dependencies are instantiated in an unspecified order along with any
// dependencies that they might have.
//
// The function may return an error to indicate failure. The error will be
// returned to the caller as-is.
func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(function, opts...)
}
