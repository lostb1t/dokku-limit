package main

import (
	"flag"
	limit "github.com/sarendsen/dokku-limit/src/limit"
	"os"
)

func main() {
	args := flag.NewFlagSet("limit:set-default", flag.ExitOnError)
	args.Parse(os.Args[2:])
	limit.CommandSetDefault(args.Args())
}
