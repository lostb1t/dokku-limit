package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dokku/dokku/plugins/network"
)

// write the ipaddress to stdout for a given app container
func main() {
	flag.Parse()
	appName := flag.Arg(0)
	procType := flag.Arg(1)
	containerID := flag.Arg(2)

	ipAddress := network.GetContainerIpaddress(appName, procType, containerID)
	fmt.Fprintln(os.Stdout, ipAddress)
}
