package config

import (
	"fmt"
	"strings"
	resource "github.com/sarendsen/dokku-limit/pkg/resource"
	"github.com/dokku/dokku/plugins/common"
)


func CommandSet(args []string, noRestart bool) error {
	if len(args) == 0 {
		common.LogFail("Please specify an app")
	}
	if len(args) == 1 {
		common.LogFail("Please specify an process")
	}
	if len(args) == 2 {
		common.LogFail("Please specify at least 1 limit")
	}
	appName := args[0]
	procName := args[1]
	verifyAppName(appName)
	new_limits := resource.Parse(args[2:])

	// Check if process exists.
	app_processes := resource.GetAppProcs(appName)
	if !app_processes[procName] {
		common.LogWarn(fmt.Sprintf("WARNING: Process \"%s\" does not exists, setting anyway.", procName))
	}
    
    // Load current resource limits or initiate new.
	limits := resource.LoadForApp(appName)
	if limits == nil {
		limits = resource.Limits{}
	}

	if limits[procName] == nil {
		limits[procName] = resource.Defaults()
	}

	// Set new limits
	for typ, limit := range new_limits {
		limits[procName][typ] = limit
	}

	limits.SaveToApp(appName)

	// todo print new limits

	if !noRestart {
		if !common.IsDeployed(appName) {
			common.LogFail("App has not been deployed, cannot restart.")
		}
		triggerRestart(appName)
	}

	return nil
}


func CommandReport(args []string) {
	apps := make(map[string]resource.Limits)

	if len(args) == 1 {
		appName := args[0]
		verifyAppName(appName)
		apps[appName] = resource.LoadForApp(appName)
	} else {
		appNames, _ := common.DokkuApps()
		for _, appName := range appNames {
			apps[appName] = resource.LoadForApp(appName)
		}
	}

	if apps == nil {
		fmt.Print("No limits set")
	}

	// todo: better looking formatting
	for appName, limits := range apps {
		for procName, resources := range limits {
			fmt.Printf("=====%s=====\n", appName)
			formatLimits(procName, resources)
		}
	}
}


func formatLimits(procName string, resources resource.Resources) {
	limits := make([]string, 0, len(resources))
	for typ, limit := range resources {
		limits = append(limits, fmt.Sprintf("%s=%s", typ, resource.FormatLimit(typ, limit)))
	} 
	fmt.Printf("%s:\t%s\n", procName, strings.Join(limits, "\t"))
}


// INTERNAL. have to move

func triggerRestart(appName string) {
	common.LogInfo1(fmt.Sprintf("Restarting app %s", appName))
	if err := common.PlugnTrigger("app-restart", appName); err != nil {
		common.LogWarn(fmt.Sprintf("Failure while restarting app: %s", err))
	}
}


func verifyAppName(appName string) {
	err := common.VerifyAppName(appName)
	if err != nil {
		common.LogFail(err.Error())
	}
}
