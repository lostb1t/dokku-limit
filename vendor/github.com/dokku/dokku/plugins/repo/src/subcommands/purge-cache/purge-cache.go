package main

import (
	"flag"
	"os"
	"strings"

	"github.com/dokku/dokku/plugins/common"
)

// deletes the contents of the build cache stored in the repository
func main() {
	flag.Parse()
	appName := flag.Arg(1)
	if appName == "" {
		common.LogFail("Please specify an app to run the command on")
	}
	err := common.VerifyAppName(appName)
	if err != nil {
		common.LogFail(err.Error())
	}

	cacheDir := strings.Join([]string{common.MustGetEnv("DOKKU_ROOT"), appName, "cache"}, "/")
	dokkuGlobalRunArgs := common.MustGetEnv("DOKKU_GLOBAL_RUN_ARGS")
	image := common.GetDeployingAppImageName(appName, "", "")
	if info, _ := os.Stat(cacheDir); info != nil && info.IsDir() {
		purgeCacheCmd := common.NewShellCmd(strings.Join([]string{"docker run --rm", dokkuGlobalRunArgs,
			"-v", strings.Join([]string{cacheDir, ":/cache"}, ""), image,
			`find /cache -depth -mindepth 1 -maxdepth 1 -exec rm -Rf {} ;`}, " "))
		purgeCacheCmd.Execute()
		err := os.MkdirAll(cacheDir, 0644)
		if err != nil {
			common.LogFail(err.Error())
		}
	}
}
