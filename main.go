package main

import (
	"fmt"
	"os"

	"github.com/ahrirpc/arpc-go/utils"
)

func print_help() {
	var text string = `Usage: arpc [options] [dir]
Options:
  -h, --help     output usage information
  -v, --version  output the version number

  -i, --input    input dir
  -o, --output   output dir
`
	fmt.Println(text)
}

func print_version() {
	fmt.Println("arpc-go version 0.0.1")
}

func main_() {
	args := os.Args
	len_args := len(args)

	switch len_args {
	case 1:
		fmt.Println("No args")
	case 2:
		if args[1] == "--help" || args[1] == "-h" {
			print_help()
		} else if args[1] == "--version" || args[1] == "-v" {
			print_version()
		} else {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", "Invalid args")
		}
	case 5:
		var input string = ""
		var output string = ""
		if args[1] == "--input" || args[1] == "-i" {
			input = args[2]
		}
		if args[3] == "--output" || args[3] == "-o" {
			output = args[4]
		}
		if input == "" || output == "" {
			fmt.Fprintf(os.Stderr, "Fatal error: %s\n", "Invalid args")
			os.Exit(1)
		}
		utils.Compiles(input, output)
	default:
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", "Invalid args")
	}
}

func main() {
	main_()
}
