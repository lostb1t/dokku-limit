package main

import (
	"fmt"
	resource "github.com/sarendsen/dokku-limit/src/resource"
	"github.com/dokku/dokku/plugins/common"
	"os"
)

func main() {
	appName := os.Args[1]
	defaults := resource.Defaults()
	var save bool = false

	limits := resource.LoadForApp(appName)
	if limits == nil {
		common.LogInfo1("No resource limit file detected, creating")
		limits = resource.Limits{}
		save = true
	}

	processes := resource.GetAppProcs(appName)
	for procName, _ := range processes {
		if _, ok := limits[procName]; !ok {
			common.LogInfo1(fmt.Sprintf("No resource limits found for process \"%s\", creating from defaults", procName))
			limits[procName] = defaults
			save = true
		}
	}

	// todo remove old procs from limits file?
	// todo print limits

	if save {
		limits.SaveToApp(appName)
	}
}
