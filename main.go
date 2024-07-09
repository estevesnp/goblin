package main

import (
	"fmt"
	"os"

	"github.com/estevesnp/dsb/pkg/repl"
)

func main() {
	fmt.Println("Starting the DSB REPL")

	repl.Start(os.Stdin, os.Stdout)
}
