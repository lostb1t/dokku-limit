package main

import (
	"os"
	"fmt"
	"strings"
	//limit "github.com/sarendsen/dokku-limit/pkg/limit"
	resource "github.com/sarendsen/dokku-limit/pkg/resource"
)

func main() {
	// $APP $IMAGE_TAG [$PROC_TYPE $CONTAINER_INDEX]
	appName := os.Args[1]
	procName := os.Args[2]

	limits := resource.LoadForApp(appName)
	if limits[procName] == nil {
		return
	}

	args := limits.DockerOptions(procName)
	if args != nil {
		fmt.Printf("%s", strings.Join(args, " "))
	}
}