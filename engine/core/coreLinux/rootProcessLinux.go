package coreLinux

import (
	"fmt"
	"os/exec"
	"strings"

	"XCalate/engine/utils"

	"github.com/projectdiscovery/gologger"
)

// CheckRootProcesses lists unique user-space processes running as root (ignores kernel threads)
func CheckRootProcesses() {
	fmt.Println()
	cmd := "ps aux | awk '{print $1,$2,$9,$10,$11}'"
	procsOutput, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		gologger.Error().Msgf("Failed to get processes: %s", err)
		return
	}

	lines := strings.Split(string(procsOutput), "\n")
	seen := make(map[string]bool)
	superUsers := []string{"root"}

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		user := fields[0]
		command := fields[4]

		// Skip kernel threads
		if strings.HasPrefix(command, "[") && strings.HasSuffix(command, "]") {
			continue
		}

		// Only show the first occurrence of each unique command
		if contains(superUsers, user) && !seen[command] {
			seen[command] = true
			gologger.Print().Label(utils.Res.String()).Msg(line)
		}
	}
}

// contains checks if a slice contains a given string
func contains(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}
