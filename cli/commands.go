package main

import (
	"github.com/jbenet/commander"
	"strings"
	"time"
	u "github.com/monesign/go-monesign/util"
)

var cmdMonesignCommands = &commander.Command{
	UsageLine: "commands",
	Short:     "List all available commands.",
	Long: `monesign commands - List all available commands.

    Lists all available commands (and sub-commands) and exits.
  `,
	Run: commandsCmd,
	Subcommands: []*commander.Command{
		cmdMonesignCommandsHelp,
	},
}
func commandsCmd(c *commander.Command, args []string) error {
	var listCmds func(c *commander.Command)
	listCmds = func(c *commander.Command) {
		u.POut("%s\n", c.FullSpacedName())
		for _, sc := range c.Subcommands {
			listCmds(sc)
		}
	}

	listCmds(c.Parent)
	return nil
}

var cmdMonesignCommandsHelp = &commander.Command{
	UsageLine: "help",
	Short:     "List all available commands' help pages.",
	Long: `monesign commands help - List all available commands's help pages.

    Shows the pages of all available commands (and sub-commands) and exits.
    Outputs a markdown document.
  `,
	Run: commandsHelpCmd,
}

func commandsHelpCmd(c *commander.Command, args []string) error {
	u.POut(referenceHeaderMsg)
	u.POut("Generated on %s.\n\n", time.Now().UTC().Format("2006-01-02"))

	var printCmds func(*commander.Command, int)
	printCmds = func(c *commander.Command, level int) {
		u.POut("%s ", strings.Repeat("#", level))
		u.POut("%s\n\n", c.FullSpacedName())
		u.POut("```\n")
		u.POut("%s\n", c.Long)
		u.POut("```\n\n")

		for _, sc := range c.Subcommands {
			printCmds(sc, level+1)
		}
	}

	printCmds(c.Parent.Parent, 1)
	return nil
}

const referenceHeaderMsg = `
# monesign command reference

This document lists every monesign command (including subcommands), along with
its help page. It can be viewed by running 'monesign commands help'.

`