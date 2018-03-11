package config

import (
	"fmt"
	"github.com/dokku/dokku/plugins/common"
	columnize "github.com/ryanuber/columnize"
	resource "github.com/sarendsen/dokku-limit/src/resource"
	"strings"
)

func CommandSet(args []string, noRestart bool) error {
	appName, procName := getCommonArgs(args)
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

	if common.IsDeployed(appName) {
		gitGcCmd := common.NewShellCmd("docker update " + strings.Join(limits[procName].DockerOptions()))
		common.LogInfo1("docker update " + strings.Join(limits[procName].DockerOptions()))
		gitGcCmd.Execute()
	}

	common.LogInfo1("Limits set")
	common.LogVerbose(formatLimits(procName, limits[procName]))

	// if !noRestart {
	// 	if !common.IsDeployed(appName) {
	// 		common.LogFail("App has not been deployed, cannot restart.")
	// 	}
	// 	triggerRestart(appName)
	// }

	return nil
}

func CommandUnSet(args []string, noRestart bool) error {
	appName, procName := getCommonArgs(args)

	types := make(map[resource.Type]bool)
	for _, typName := range args[2:] {
		typ, ok := resource.ToType(typName)
		if ok {
			types[typ] = false
		}
	}

	limits := resource.LoadForApp(appName)
	if limits == nil {
		common.LogInfo1(fmt.Sprintf("No limits set for \"%s\"", appName))
		return nil
	}

	resources := limits[procName]
	if resources == nil {
		common.LogInfo1(fmt.Sprintf("No limits set for \"%s\"", procName))
		return nil
	}

	// Unset limits
	var restart bool = false
	for typ, _ := range types {
		if _, ok := resources[typ]; ok {
			common.LogInfo1(fmt.Sprintf("Unsetting \"%s\"", typ))
			delete(resources, typ)
			restart = true
		}
	}

	limits.SaveToApp(appName)

	common.LogInfo1("Limits unset")
	common.LogVerbose(formatLimits(procName, resources))

	if !noRestart && restart {
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
		fmt.Println("No limits set")
	}

	for appName, limits := range apps {
		for procName, resources := range limits {
			common.LogInfo2(appName + " limits")
			fmt.Println(formatLimitsTable(procName, resources))
		}
	}
}

// Helpers

func formatLimits(procName string, resources resource.Resources) string {
	limits := make([]string, 0, len(resource.ResourceTypes))

	for _, typ := range resource.ResourceTypes {
		if r, ok := resources[typ]; ok {
			limits = append(limits, fmt.Sprint(string(typ), ": ", resource.Format(typ, r)))
		} else {
			limits = append(limits, fmt.Sprint(string(typ), ": ", "-"))
		}
	}

	return strings.Join(limits, " ")
}


func formatLimitsTable(procName string, resources resource.Resources) string {
	config := columnize.DefaultConfig()
	config.Delim = "|"
	config.Empty = "-"

	header := make([]string, 0, len(resource.ResourceTypes))
	limits := make([]string, 0, len(resource.ResourceTypes))
	for _, typ := range resource.ResourceTypes {
		header = append(header, string(typ))
		if r, ok := resources[typ]; ok {
			limits = append(limits, resource.Format(typ, r))
		} else {
			limits = append(limits, "")
		}
	}

	content := []string{
		strings.Join(header, config.Delim),
		strings.Join(limits, config.Delim),
	}
	return columnize.Format(content, config)
}


func triggerRestart(appName string) {
	common.LogInfo1(fmt.Sprintf("Restarting app %s", appName))
	if err := common.PlugnTrigger("app-restart", appName); err != nil {
		common.LogWarn(fmt.Sprintf("Failure while restarting app: %s", err))
	}
}

func getCommonArgs(args []string) (appName string, procName string) {
	if len(args) == 0 {
		common.LogFail("Please specify an app")
	}

	appName = args[0]
	verifyAppName(appName)

	if len(args) == 1 {
		common.LogFail("Please specify an process")
	}
	if len(args) == 2 {
		common.LogFail("Please specify at least 1 resource")
	}
	procName = args[1]

	return appName, procName
}

func verifyAppName(appName string) {
	err := common.VerifyAppName(appName)
	if err != nil {
		common.LogFail(err.Error())
	}
}
