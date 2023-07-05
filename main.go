package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/achimwinter/monkey-language/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is Monkey Language!\n", user.Username);
	fmt.Printf("Enter any command!\n")
	repl.Start(os.Stdin, os.Stdout)
}