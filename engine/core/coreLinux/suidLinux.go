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

	// Only log error if there's no output at all
	if err != nil && len(output) == 0 {
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

func CheckCapability() {
	cmd := exec.Command("sh", "-c", "getcap -r / 2>/dev/null")
	output, err := cmd.CombinedOutput()

	// If command errored and produced no output, log and return
	if err != nil && len(output) == 0 {
		gologger.Error().Msgf("Error while searching for capabilities: %v", err)
		return
	}

	// If no output at all, likely getcap not present or nothing found anywhere
	if len(output) == 0 || strings.TrimSpace(string(output)) == "" {
		gologger.Print().Label(utils.Sad.String()).Msg("No capabilities binary")
		return
	}

	// Collect lines that contain cap_setuid+ep
	var found []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.Contains(line, "cap_setuid+ep") {
			found = append(found, line)
		}
	}

	// Print results based on whether any cap_setuid+ep entries were found
	if len(found) > 0 {
		gologger.Print().Label(utils.Res.String()).Msg("`cap_setuid+ep` binaries found:")
		for _, l := range found {
			gologger.Info().Msg(l)
		}
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("No `cap_setuid+ep` binaries found")
	}
}
