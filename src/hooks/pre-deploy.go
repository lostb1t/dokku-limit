package main

import (
	"fmt"
	limit "github.com/sarendsen/dokku-limit/src/limit"
	resource "github.com/sarendsen/dokku-limit/src/resource"
	"github.com/dokku/dokku/plugins/common"
	"os"
)

func main() {
	appName := os.Args[1]
	imageTag := os.Args[2]
	defaults := resource.Defaults()
	var save bool = false

	// sucks but need this because our plugin runs before ps because of the name...
	// we need the SCALE file.
	dokkuCmd := common.NewShellCmd("extract_procfile " + appName + " " + imageTag)
	dokkuCmd.ShowOutput = false
	dokkuCmd.Execute()
	dokkuCmd = common.NewShellCmd("generate_scale_file " + appName + " " + imageTag)
	dokkuCmd.ShowOutput = false
	dokkuCmd.Execute()

	limits := resource.LoadForApp(appName)
	if limits == nil {
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

	common.LogInfo1(fmt.Sprintf("Resource limits"))
	for procName, _ := range processes {
		common.LogInfo1(fmt.Sprintf("%s %s", procName, limit.FormatLimits(limits[procName])))
	}

	if save {
		limits.SaveToApp(appName)
	}
}
