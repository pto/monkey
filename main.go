// Monkey starts a REPL for the Monkey programming language.
package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/pto/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	name := user.Name
	if name == "" {
		name = user.Username
	}
	fmt.Printf("Hello, %s! This is the Monkey programming language.\n", name)
	repl.Start(os.Stdin, os.Stdout)
}
