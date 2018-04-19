package command_pattern

type Command interface {
	do() string
}

type CommandFunc (func() string)

func (f CommandFunc) do() string {
	return f()
}

func run() string {
	return "Run"
}

func walk() string {
	return "Walk"
}
