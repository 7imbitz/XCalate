package coreLinux

import (
	"XCalate/engine/utils"
	"fmt"
	"os/exec"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func CheckKernel() {
	// CVE-2022-0847
	checkCve220847()
	// CVE-2021-3493
	checkCve213493()
	fmt.Println()
}

func checkCve220847() {
	cmd := exec.Command("sh", "-c", "uname -a | awk '{print $3}'")
	output, err := cmd.Output()
	if err != nil {
		gologger.Error().Msgf("Error running command: %s", err)
		return
	}

	kernelVersion := strings.TrimSpace(string(output))
	gologger.Info().Msgf("Kernel Version: %s", kernelVersion)

	vulnerableRangeStart := "5.8.0"
	vulnerableRangeEnd := "5.16.11"
	patchVersions := []string{"5.16.11", "5.15.25", "5.10.102"}

	if isVulnerable(kernelVersion, vulnerableRangeStart, vulnerableRangeEnd, patchVersions) {
		gologger.Print().Label(utils.Res.String()).Msg("Kernel version MIGHT be vulnerable to CVE-2022-0847")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("Kernel version was either not vulnerable or patched for CVE-2022-0847.")
	}
	fmt.Println()
	gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `uname -a | awk '{print $3}'`")
}

func checkCve213493() {
	vulnerableVersions := []string{
		"Ubuntu 20.10",
		"Ubuntu 20.04 LTS",
		"Ubuntu 19.04",
		"Ubuntu 18.04 LTS",
		"Ubuntu 16.04 LTS",
		"Ubuntu 14.04 ESM",
	}

	cmd := exec.Command("sh", "-c", "cat /etc/issue")
	output, err := cmd.Output()
	if err != nil {
		gologger.Error().Msgf("Error running command: %s", err)
		return
	}

	osVersion := strings.TrimSpace(string(output))
	gologger.Info().Msgf("Name and OS Version: %s", osVersion)

	// Check if the output matches any of the vulnerable versions
	for _, version := range vulnerableVersions {
		if strings.Contains(osVersion, version) {
			gologger.Print().Label(utils.Res.String()).Msg("Kernel version MIGHT be vulnerable to CVE-2021-3493")
			return
		}
	}
	gologger.Print().Label(utils.Sad.String()).Msg("Kernel version was either not vulnerable or patched for CVE-2021-3493.")

}

func isVulnerable(version, rangeStart, rangeEnd string, patchVersions []string) bool {
	if compareVersions(version, rangeStart) >= 0 && compareVersions(version, rangeEnd) < 0 {
		return true
	}

	for _, patchVersion := range patchVersions {
		if version == patchVersion {
			return false // It's a patch version
		}
	}

	return false
}

func compareVersions(a, b string) int {
	// Split version strings into arrays of integers
	aParts := splitVersion(a)
	bParts := splitVersion(b)

	// Compare each part of the version
	for i := 0; i < len(aParts) && i < len(bParts); i++ {
		if aParts[i] > bParts[i] {
			return 1
		} else if aParts[i] < bParts[i] {
			return -1
		}
	}

	// If all parts are equal, compare the length
	return len(aParts) - len(bParts)
}

func splitVersion(version string) []int {
	var parts []int
	for _, part := range strings.Split(version, ".") {
		var num int
		_, err := fmt.Sscanf(part, "%d", &num)
		if err == nil {
			parts = append(parts, num)
		}
	}
	return parts
}
