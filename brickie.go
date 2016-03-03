package main

import (
	"log"

	"github.com/ivanfoo/brickie/commands"

	"github.com/jessevdk/go-flags"
)

func main() {
	parser := flags.NewNamedParser("brickie", flags.Default)
	parser.AddCommand("build", "build images", "", commands.NewCmdBuild())
	parser.AddCommand("push", "push images", "", commands.NewCmdPush())

	_, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}
}
