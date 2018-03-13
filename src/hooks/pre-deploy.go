package main

import (
	"fmt"
	"strings"
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
	// we need the updated SCALE file.
	// todo: create a pull request for plugn system supporting plugin weights
	pluginRoot := common.MustGetEnv("PLUGIN_PATH")
	psFile := strings.Join([]string{pluginRoot, "enabled", "ps", "pre-deploy"}, "/")
	dokkuCmd := common.NewShellCmd(fmt.Sprintf("%s %s %s", psFile, appName, imageTag))
	dokkuCmd.ShowOutput = false
	dokkuCmd.Execute()

	limits := resource.LoadForApp(appName)
	if limits == nil {
		limits = resource.Limits{}
		save = true
	}

	// check if all procs have limits
	processes := resource.GetAppProcs(appName)
	for procName, _ := range processes {
		if _, ok := limits[procName]; !ok {
			common.LogInfo1(fmt.Sprintf("No resource limits found for process \"%s\", creating from defaults", procName))
			limits[procName] = defaults
			save = true
		}
	}

	// remove non existing procs
	for procName, _ := range limits {
		if _, ok := processes[procName]; !ok {
			delete(limits, procName)
			save = true
		}
	}


	for procName, _ := range processes {
		common.LogInfo1(fmt.Sprintf("%s limits: [%s]", procName, limit.FormatLimits(limits[procName])))
	}

	if save {
		limits.SaveToApp(appName)
	}
}
