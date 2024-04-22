package main

import (
	"flag"
	"os"

	"github.com/bobmaertz/tcli/internal/version"
)

var commands = map[string]func([]string){
	"version": version.Print,
}

func init() {
	//TODO:  Define flags
	flag.Parse()
}

func main() {
	args := flag.Args()

	if len(args) < 1 {
		usage()
	}

	// Take the first argument as the command to run with the second argument as the argument to the command
	commands[args[0]](args[1:])
}

func usage() {
	flag.Usage()
	os.Exit(0)
}
