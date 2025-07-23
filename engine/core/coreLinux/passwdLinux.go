package coreLinux

import (
	"XCalate/engine/utils"

	"github.com/projectdiscovery/gologger"
)

// Writable /etc/passwd
func CheckPasswdWritable() {

	passwdWritable, errPasswdWritable := utils.IsWorldWritable(utils.PasswdPath)
	if errPasswdWritable != nil {
		gologger.Error().Msgf("Error checking world writability: %s", errPasswdWritable)
		return
	}
	if passwdWritable {
		gologger.Print().Label(utils.Res.String()).Msg("Passwd file is world-writable")
		gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `ls -lat /etc/passwd`")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Passwd file is not world-writable")
	}
}
