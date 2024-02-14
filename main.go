package main

import (
	"fmt"
	"io"
	"os"
	"os/user"

	"github.com/baxromumarov/monkey-programming-language/repl"
)


const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func main() {
	userData, err := user.Current()
	if err != nil {
		panic(err)
	}
	io.WriteString(os.Stdin, MONKEY_FACE)

	fmt.Printf("Hello %s! This is the monkey-programming-language programming language!\n", userData.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
