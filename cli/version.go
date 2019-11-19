package main

import (
	"github.com/jbenet/commander"
	u "github.com/monesign/go-monesign/util"
)

const Version = "0.1.0"

var cmdMonesignVersion = &commander.Command{
	UsageLine: "version",
	Short:     "Show monesign version information.",
	Long: `monesign version - Show monesign version information.

    Returns the current version of monesign and exits.
  `,
	Run:  versionCmd,
}

func init() {
	cmdMonesignVersion.Flag.Bool("number", false, "show only the number")
}

func versionCmd(c *commander.Command, _ []string) error {
	number := c.Flag.Lookup("number").Value.Get().(bool)
	if !number {
		u.POut("monesign version ")
	}
	u.POut("%s\n", Version)
	return nil
}