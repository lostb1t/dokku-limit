package main

import (
	"os"
	"fmt"
	"strings"
	resource "github.com/sarendsen/dokku-limit/src/resource"
)

func main() {
	// $CALLER $APP $IMAGE_TAG [$PROC_TYPE $CONTAINER_INDEX]
	// Seems some calls don't include the proctype
	if len(os.Args) < 4 {
		return
	}

	appName := os.Args[1]
	procName := os.Args[3]

	limits := resource.LoadForApp(appName)
	if limits == nil {
		return
	}

	if limits[procName] == nil {
		return
	}

	args := limits.DockerOptions(procName)
	if args != nil {
		fmt.Printf(" %s", strings.Join(args, " "))
	}
}