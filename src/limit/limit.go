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

	// make sure all resource types are set
	resource.SetDefaults(limits[procName])

	limits.SaveToApp(appName)

	common.LogInfo1("Limits set")
	common.LogVerbose(FormatLimits(limits[procName]))

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
		limits := resource.LoadForApp(appName)
		if limits != nil {
			apps[appName] = resource.LoadForApp(appName)
		}
	} else {
		appNames, _ := common.DokkuApps()
		for _, appName := range appNames {
			limits := resource.LoadForApp(appName)
			if limits != nil {
				apps[appName] = resource.LoadForApp(appName)
			}
		}
	}

	if len(apps) == 0 {
		fmt.Println("No limits set")
	}

	for appName, limits := range apps {
		for procName, resources := range limits {
			resource.SetDefaults(limits[procName])
			common.LogInfo2(appName + " " + procName + " limits")
			fmt.Println(formatLimitsTable(resources))
		}
	}
}


func CommandReportDefault(args []string) {
	defaultResources := resource.Defaults()
	common.LogInfo2("Default limits")
	fmt.Println(formatLimitsTable(defaultResources))
}


func CommandSetDefault(args []string) error {
	if len(args) == 0 {
		common.LogFail("Please specify at least 1 resource")
	}

	resources := resource.Parse(args)

	// Load current defaults
	defaultResources := resource.Defaults()

	// Set new defaults
	for typ, resource := range resources {
		defaultResources[typ] = resource
	}

	// save new defaults
	err := resource.SaveDefaults(defaultResources)
	if err != nil {
		common.LogFail(err.Error())
	}

	common.LogInfo1("Default limits saved")

	return nil
}


// Helpers

func FormatLimits(resources resource.Resources) string {
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


func formatLimitsTable(resources resource.Resources) string {
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
