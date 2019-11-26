package main

import (
	"fmt"
	"github.com/gonuts/flag"
	"github.com/jbenet/commander"
	u "github.com/monesign/go-monesign/util"
	"os"
)

var CmdMonesign = &commander.Command{
	UsageLine: "monesign [<flags>] <command> [<args>]",
	Short:     "global versioned p2p merkledag file system",
	Long: `monesign - global versioned p2p merkledag file system

Basic commands:

    add <path>    Add an object to monesign.
    cat <ref>     Show monesign object data.
    ls <ref>      List links from an object.
    refs <ref>    List link hashes from an object.

Tool commands:

    config      Manage configuration.
    version     Show monesign version information.
    commands    List all available commands.

Advanced Commands:

    mount       Mount a monesign a read-only mountpoint.

Use "monesign help <command>" for more information about a command.
`,
	Run: monesignCmd,
	Subcommands: []*commander.Command{
		cmdMonesignVersion,
		// cmdMonesignConfig,
		cmdMonesignCommands,
	},
	Flag: *flag.NewFlagSet("monesign", flag.ExitOnError),
}

func monesignCmd(c *commander.Command, args []string) error {
	u.POut(c.Long)
	return nil
}

func main(){
	err := CmdMonesign.Dispatch(os.Args[:1])
	if err != nil {
		if len(err.Error()) > 0 {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		os.Exit(1)
	}
	return
}