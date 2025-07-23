package coreLinux

import (
	"os/exec"
	"strings"

	"XCalate/engine/utils"

	"github.com/projectdiscovery/gologger"
)

func CheckRootProcesses() {
	procsOutput, err := exec.Command("sh", "-c", "ps aux | awk '{print $1,$2,$9,$10,$11}'").Output()
	if err != nil {
		gologger.Error().Msgf("Failed to get processes: %s", err)
		return
	}

	pkgsOutput, err := exec.Command("sh", "-c", "dpkg -l | awk '{$1=$4=\"\"; print $0}'").Output()
	if err != nil {
		gologger.Error().Msgf("Failed to get packages: %s", err)
		return
	}

	superUsers := []string{"root"} // add more if needed
	procLines := strings.Split(string(procsOutput), "\n")
	pkgLines := strings.Split(string(pkgsOutput), "\n")

	procDict := make(map[string][]string)

	for _, proc := range procLines {
		for _, user := range superUsers {
			if user != "" && strings.Contains(proc, user) {
				fields := strings.Fields(proc)
				if len(fields) < 5 {
					continue
				}
				procName := fields[4]
				if strings.Contains(procName, "/") {
					parts := strings.Split(procName, "/")
					procName = parts[len(parts)-1]
				}
				if len(procName) < 3 {
					continue
				}
				relatedPkgs := procDict[proc]
				for _, pkg := range pkgLines {
					if strings.Contains(pkg, procName) && !contains(relatedPkgs, pkg) {
						relatedPkgs = append(relatedPkgs, pkg)
					}
				}
				procDict[proc] = relatedPkgs
			}
		}
	}

	for key, val := range procDict {
		gologger.Print().Label(utils.Res.String()).Msg(key)
		if len(val) > 0 && val[0] != "" {
			gologger.Print().Label(utils.Bsh.String()).Msg("Possible Related Packages:")
			for _, pkg := range val {
				gologger.Info().Msg(" " + pkg)
			}
		}
	}
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
