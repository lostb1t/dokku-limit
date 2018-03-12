package main

import (
	"bufio"
	"fmt"
	resource "github.com/sarendsen/dokku-limit/src/resource"
	"os"
	"strings"
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
		limits = resource.Limits{
			procName: resource.Defaults(),
		}
	}

	resources, ok := limits[procName]
	if !ok {
		limits[procName] = resource.Defaults()
	}

	args := resources.DockerOptions()
	if args != nil {
		fmt.Print(stdin, strings.Join(args, " "))
	} else {
		fmt.Printf("%s", stdin)
	}
}
