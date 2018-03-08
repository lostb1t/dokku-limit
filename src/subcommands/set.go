package main

import (
	"flag"
	limit "github.com/sarendsen/dokku-limit/src/limit"
	"os"
)

// set the given entries to the specified environment
func main() {
	args := flag.NewFlagSet("limit:set", flag.ExitOnError)
	noRestart := args.Bool("no-restart", false, "--no-restart: no restart")
	args.Parse(os.Args[2:])
	limit.CommandSet(args.Args(), *noRestart)
}
