package commandpattern

// Command pattern interface.
type Command interface {
	do() string
}

// ----------------------------------------------------------------------------

// CommandFunc makes function to command pattern object.
type CommandFunc (func() string)

func (f CommandFunc) do() string {
	return f()
}

// ----------------------------------------------------------------------------

func run() string {
	return "Run"
}

func walk() string {
	return "Walk"
}
