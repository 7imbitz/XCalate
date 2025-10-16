package main

import (
	"XCalate/engine/core"
	"XCalate/engine/report"
	"XCalate/engine/utils"
	"runtime"

	"github.com/projectdiscovery/gologger"
)

func main() {
	utils.ShowBanner()
	os := runtime.GOOS
	switch os {
	case "linux":
		report.MarkTask("Task 1 - Linux Checks", true)
		gologger.Info().Msg("OS is linux, running for linux PE.")
		core.LinPrivEscChecker()
	case "window":
		gologger.Info().Msg("OS is windows, running for window PE.")
		//TODO
	case "darwin":
		gologger.Fatal().Msg("Running on Mac")
	}

}
