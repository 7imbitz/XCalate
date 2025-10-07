package coreLinux

import (
	"XCalate/engine/utils"
	"fmt"

	"github.com/projectdiscovery/gologger"
)

// Writable /etc/passwd
func CheckPasswdWritable() {
	gologger.Info().Msg("Checking if Passwd file is world-writable ↓")

	passwdWritable, errPasswdWritable := utils.IsWorldWritable(utils.PasswdPath)
	if errPasswdWritable != nil {
		gologger.Error().Msgf("Error checking world writability: %s", errPasswdWritable)
		return
	}
	if passwdWritable {
		gologger.Print().Label(utils.Res.String()).Msg("Passwd file is world-writable")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Passwd file is not world-writable")
	}
	fmt.Println()
	gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `ls -lat /etc/passwd`")
}
