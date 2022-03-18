package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s, welcome to Monkey\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
