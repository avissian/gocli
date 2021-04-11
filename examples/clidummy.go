/*
clidummy demonstrates the basics of using the gocli library
*/
package main

import (
	"fmt"

	"github.com/avissian/gocli"
)

func main() {

	cli := gocli.MkCLI("Welcome to this dummy CLI.")

	// register help Option with cli.Help as callback
	cli.AddOption("help", "prints this help message", cli.Help)

	// register rather long command as Option with custom callback function
	cli.AddOption("kapitänsmützenabzeichen", "just an example of a long cmd name not breaking the help formatting", func(_ []string) string { return cli.Exit([]string{"fnord"}) })

	// Options with empty help string registered as hidden. Should not appear in "help" list
	cli.AddOption("quit", "", cli.Exit)
	cli.AddOption("?", "", cli.Help)

	// demonstrate argument passing
	cli.AddOption("quote", "qotes all arguments", func(args []string) string {
		res := ""
		for i, arg := range args[1:] {
			if i != 0 {
				res += " "
			}
			res += "\"" + arg + "\""
		}
		return res
	})

	// other demonstration argument passing
	for rune := 'a'; rune <= 'c'; rune++ {
		str := string(rune)
		cli.AddOption(str, fmt.Sprintf("Writes %s on screen", str), func(args []string) string {
			message := fmt.Sprintf("Pressed %s key", args[0])
			if len(args) > 1 {
				message += fmt.Sprintf(" with arguments %v", args[1:])
			}
			return message
		},
		)
	}

	cli.AddSeparator()

	// register exit Option with cli.Exit as callback
	cli.AddOption("exit", "exits the input loop", cli.Exit)

	// default callback
	cli.DefaultOption(func(args []string) string {
		return fmt.Sprintf("%s: command not found, type 'help' or '?' for help", args[0])
	})

	// run the main loop
	cli.Loop("dummyprompt> ")

	// after breaking cli.Loop
	fmt.Println("(this part of the code is only reached when the cli loop returns)")

}
