package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	columnize "github.com/ryanuber/columnize"
	limit "github.com/sarendsen/dokku-limit/src/limit"
)

const (
	helpHeader = `Usage: dokku limit [<app>|--global]
Display all global or app-specific config vars
Additional commands:`

	helpContent = `
    limit (<app>), Pretty-print limits
    limit:set <app> <proc> [memory=VALUE cpu=VALUE] [--no-restart], Set one or more limits for app/process pair
    limit:unset <app> <proc> [memory cpu] [--no-restart], Unset one or more limits for app/process pair
    limit:set-default [memory=VALUE cpu=VALUE], Set default resources. App(s) need to be restarted to take effect.
`
)

func main() {
	flag.Usage = usage
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "limit", "limit:report":
		args := flag.NewFlagSet("limit:report", flag.ExitOnError)
		args.Parse(os.Args[2:])
		limit.CommandReport(args.Args())
	case "help":
		fmt.Print("\n    limit, App resource management.\n")
	case "limit:help":
		usage()
	default:
		dokkuNotImplementExitCode, err := strconv.Atoi(os.Getenv("DOKKU_NOT_IMPLEMENTED_EXIT"))
		if err != nil {
			fmt.Println("failed to retrieve DOKKU_NOT_IMPLEMENTED_EXIT environment variable")
			dokkuNotImplementExitCode = 10
		}
		os.Exit(dokkuNotImplementExitCode)
	}
}

func usage() {
	config := columnize.DefaultConfig()
	config.Delim = ","
	config.Prefix = "    "
	config.Empty = ""
	content := strings.Split(helpContent, "\n")[1:]
	fmt.Println(helpHeader)
	fmt.Println(columnize.Format(content, config))
}
