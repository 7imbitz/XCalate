package coreLinux

import (
	"XCalate/engine/utils"
	"os"

	"github.com/projectdiscovery/gologger"
)

// CheckNFS checks if /etc/exports file exists, indicating potential NFS export configuration
func CheckNFS() {
	// Check if /etc/exports file exists
	if _, err := os.Stat("/etc/exports"); os.IsNotExist(err) {
		gologger.Print().Label(utils.Sad.String()).Msg("/etc/exports file not found")
		return
	} else if err != nil {
		// Handle other possible errors
		gologger.Error().Msgf("Error checking /etc/exports: %s", err)
		return
	}

	// File exists
	gologger.Print().Label(utils.Res.String()).Msg("Found NFS export configuration! `/etc/exports`")
	gologger.Info().Msg("Files created via NFS inherit the remote user's ID. If the user is root, and root squashing is enabled, the ID will instead be set to the \"nobody\" user.")
}
