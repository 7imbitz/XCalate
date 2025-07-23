package coreLinux

import (
	"XCalate/engine/utils"
	"os/exec"
	"strings"

	"github.com/projectdiscovery/gologger"
)

// CheckSUIDExec scans for SUID binaries on the system
func CheckSUIDExec() {
	cmd := exec.Command("sh", "-c", "find / -type f -perm -04000 -ls 2>/dev/null")
	output, err := cmd.CombinedOutput()
	if err != nil {
		gologger.Error().Msgf("Error while searching for SUID binaries: %s", err)
		return
	}

	if len(output) == 0 {
		gologger.Print().Label(utils.Sad.String()).Msg("No SUID binaries found")
		return
	}

	gologger.Print().Label(utils.Res.String()).Msg("SUID binaries found:")
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			gologger.Info().Msg(line)
		}
	}
}
