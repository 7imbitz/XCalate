package main

import (
	"XCalate/engine/core"
	"XCalate/engine/utils"
	"runtime"

	"github.com/projectdiscovery/gologger"
)

func main() {
	utils.ShowBanner()
	os := runtime.GOOS
	switch os {
	case "linux":
		gologger.Info().Msg("OS is linux, running for linux PE.")
		core.LinPrivEscChecker()
	case "window":
		gologger.Info().Msg("OS is windows, running for window PE.")
		//TODO 2025
	case "darwin":
		gologger.Fatal().Msg("Running on Mac")
	}

}
