package coreLinux

import (
	"XCalate/engine/utils"
	"os/exec"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// CheckNFS checks if NFS shares are configured via /etc/exports
func CheckNFS() {
	cmd := exec.Command("sh", "-c", "cat /etc/exports")
	output, err := cmd.CombinedOutput()

	// Check if file read failed
	if err != nil {
		if strings.Contains(err.Error(), "no such file") {
			gologger.Print().Label(utils.Sad.String()).Msg("/etc/exports file not found")
		} else {
			gologger.Error().Msgf("Error reading /etc/exports: %s", err)
		}
		return
	}

	trimmed := strings.TrimSpace(string(output))
	if trimmed == "" {
		gologger.Print().Label(utils.Sad.String()).Msg("/etc/exports is empty")
		return
	}

	gologger.Print().Label(utils.Res.String()).Msg("Found NFS export configuration!")
	gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `cat /etc/exports`")

	lines := strings.Split(trimmed, "\n")
	for _, line := range lines {
		gologger.Info().Msg(line)
	}
}
