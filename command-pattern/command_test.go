package commandpattern

import (
	"fmt"
)

func ExampleCommand() {
	var cmd Command

	cmd = CommandFunc(run)
	fmt.Println(cmd.do())

	cmd = CommandFunc(walk)
	fmt.Println(cmd.do())

	// Output:
	// Run
	// Walk
}
