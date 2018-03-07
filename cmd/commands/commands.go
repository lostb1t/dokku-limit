package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	limit "github.com/sarendsen/dokku-limit/pkg/limit"
	columnize "github.com/ryanuber/columnize"
)

const (
	helpHeader = `Usage: dokku limit [<app>|--global]
Display all global or app-specific config vars
Additional commands:`

	helpContent = `
    limit (<app>), Pretty-print limits
    limit:set <app> <proc> [memory=VALUE cpu=VALUE] [--no-restart], Set one or more limits for app/process pair
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