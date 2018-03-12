package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	resource "github.com/sarendsen/dokku-limit/src/resource"
)

func main() {
	// $CALLER $APP $IMAGE_TAG [$PROC_TYPE $CONTAINER_INDEX]
	reader := bufio.NewReader(os.Stdin)
	stdin, _ := reader.ReadString('\n')
	stdin = strings.TrimSuffix(stdin, "\n")

	if len(os.Args) < 4 {
		fmt.Printf("%s", stdin)
		return
	}

	appName := os.Args[1]
	procName := os.Args[3]

	limits := resource.LoadForApp(appName)
	if limits == nil {
		fmt.Printf("%s", stdin)
		return
	}

	resources, ok := limits[procName]
	if !ok {
		fmt.Printf("%s", stdin)
		return
	}

	args := resources.DockerOptions()
	if args != nil {
		fmt.Print(stdin, strings.Join(args, " "))
	} else {
		fmt.Printf("%s", stdin)
	}
	return
}