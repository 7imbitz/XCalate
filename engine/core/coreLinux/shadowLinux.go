package coreLinux

import (
	"XCalate/engine/utils"
	"fmt"

	"github.com/projectdiscovery/gologger"
)

// CheckShadowPermissions checks both world-readability and world-writability of /etc/shadow
func CheckShadowPermissions() {
	gologger.Info().Msg("Checking if Shadow file is world-readable ↓")
	readable, errReadable := utils.IsWorldReadable(utils.ShadowPath)

	if errReadable != nil {
		gologger.Error().Msgf("Error checking world readability: %s", errReadable)
	} else if readable {
		gologger.Print().Label(utils.Res.String()).Msg("Shadow file is world-readable")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Shadow file is not world-readable")
	}
	fmt.Println()

	gologger.Info().Msg("Checking if Shadow file is world-writable ↓")
	writable, errWritable := utils.IsWorldWritable(utils.ShadowPath)

	if errWritable != nil {
		gologger.Error().Msgf("Error checking world writability: %s", errWritable)
	} else if writable {
		gologger.Print().Label(utils.Res.String()).Msg("Shadow file is world-writable")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Shadow file is not world-writable")
	}
	fmt.Println()

	// Show the verification command once
	gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `ls -lat /etc/shadow`")

}
