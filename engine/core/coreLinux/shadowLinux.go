package coreLinux

import (
	"XCalate/engine/utils"

	"github.com/projectdiscovery/gologger"
)

// Readable /etc/shadow
func CheckShadowReadable() {

	shadowReadable, errShadowReadable := utils.IsWorldReadable(utils.ShadowPath)
	if errShadowReadable != nil {
		gologger.Error().Msgf("Error checking world readability: %s", errShadowReadable)
		return
	}
	if shadowReadable {
		gologger.Print().Label(utils.Res.String()).Msg("Shadow file is world-readable")
		gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `ls -lat /etc/shadow`")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Shadow file is not world-readable")
	}
}

// Writable /etc/shadow
func CheckShadowWritable() {

	shadowWritable, errShadowWritable := utils.IsWorldWritable(utils.ShadowPath)
	if errShadowWritable != nil {
		gologger.Error().Msgf("Error checking world writability: %s", errShadowWritable)
		return
	}
	if shadowWritable {
		gologger.Print().Label(utils.Res.String()).Msg("Shadow file is world-writable")
		gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `ls -lat /etc/shadow`")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Shadow file is not world-writable")
	}
}
