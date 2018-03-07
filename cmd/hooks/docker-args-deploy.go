package main

import (
	//"os"
	"fmt"
	limit "github.com/sarendsen/dokku-limit/pkg/limit"
	resource "github.com/sarendsen/dokku-limit/pkg/resource"
)

func main() {
	// $APP $IMAGE_TAG [$PROC_TYPE $CONTAINER_INDEX]
	appName := os.Args[0]
	procName := os.Args[2]

	limits := resource.LoadForApp(appName)
	args, err := limits.FormatDocker(procName)
	if err != nil {
		common.LogWarn(err.Error())
	}

	if args != nil {
		fmt.Printf("%s", strings.Join(args, " "))
	}
}