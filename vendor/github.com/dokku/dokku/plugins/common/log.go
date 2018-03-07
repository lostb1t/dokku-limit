package common

import (
	"fmt"
	"os"
)

// LogFail is the failure log formatter
// prints text to stderr and exits with status 1
func LogFail(text string) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf("FAILED: %s", text))
	os.Exit(1)
}

// LogInfo1 is the info1 header formatter
func LogInfo1(text string) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("-----> %s", text))
}

// LogInfo1Quiet is the info1 header formatter (with quiet option)
func LogInfo1Quiet(text string) {
	if os.Getenv("DOKKU_QUIET_OUTPUT") != "" {
		LogInfo1(text)
	}
}

// LogInfo2 is the info2 header formatter
func LogInfo2(text string) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("=====> %s", text))
}

// LogInfo2Quiet is the info2 header formatter (with quiet option)
func LogInfo2Quiet(text string) {
	if os.Getenv("DOKKU_QUIET_OUTPUT") == "" {
		LogInfo2(text)
	}
}

// LogVerbose is the verbose log formatter
// prints indented text to stdout
func LogVerbose(text string) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf("       %s", text))
}

// LogVerboseQuiet is the verbose log formatter
// prints indented text to stdout (with quiet option)
func LogVerboseQuiet(text string) {
	if os.Getenv("DOKKU_QUIET_OUTPUT") != "" {
		LogVerbose(text)
	}
}

// LogWarn is the warning log formatter
func LogWarn(text string) {
	fmt.Fprintln(os.Stderr, fmt.Sprintf(" !     %s", text))
}
