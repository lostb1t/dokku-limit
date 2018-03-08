package main

import (
	"flag"
	limit "github.com/sarendsen/dokku-limit/src/limit"
	"os"
)

func main() {
	args := flag.NewFlagSet("limit:unset", flag.ExitOnError)
	noRestart := args.Bool("no-restart", false, "--no-restart: no restart")
	args.Parse(os.Args[2:])
	limit.CommandUnSet(args.Args(), *noRestart)
}
