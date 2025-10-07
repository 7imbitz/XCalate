package coreLinux

import (
	"XCalate/engine/utils"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func CheckCronJobDetails() {
	gologger.Info().Msg("Checking cronjobs environment PATH ↓")

	currentUser, err := user.Current()
	if err != nil {
		gologger.Error().Msgf("Error getting current user: %s", err)
		return
	}

	content, err := os.ReadFile(utils.CronPath)
	if err != nil {
		gologger.Error().Msgf("Error reading /etc/crontab: %s", err)
		return
	}
	crontabString := string(content)

	pathLine := ""
	for _, line := range strings.Split(crontabString, "\n") {
		if strings.HasPrefix(line, "PATH=") {
			pathLine = line
			break
		}
	}
	if strings.Contains(pathLine, currentUser.HomeDir) {
		gologger.Print().Label(utils.Res.String()).Msg("PATH in /etc/crontab has home user directory")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("PATH in /etc/crontab does not have home user directory")
	}
	fmt.Println()
	gologger.Info().Msg("Checking cronjobs being run with a wildcard (*) ↓")
	gologger.Print().Label(utils.Res.String()).Msg("Content of /etc/crontab")
	lines := strings.Split(crontabString, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		fmt.Println(line)
	}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "run-parts") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		//scriptField := fields[5]
		scriptField := strings.Join(fields[6:], " ")

		// If full path
		var allScriptPaths []string

		if strings.Contains(scriptField, "cd ") && strings.Contains(scriptField, "&&") {
			parts := strings.Split(scriptField, "&&")
			if len(parts) == 2 {
				dirPart := strings.TrimSpace(strings.TrimPrefix(parts[0], "cd"))
				dirPart = strings.TrimSpace(dirPart)
				cmdPart := strings.TrimSpace(parts[1])

				if strings.Contains(cmdPart, "bash") || strings.Contains(cmdPart, "sh") {
					cmdFields := strings.Fields(cmdPart)
					if len(cmdFields) > 1 {
						scriptName := strings.TrimSpace(cmdFields[len(cmdFields)-1])
						fullPath := filepath.Join(dirPart, scriptName)
						allScriptPaths = append(allScriptPaths, fullPath)
					}
				}
			}
		} else if strings.HasPrefix(scriptField, "/") {
			// Direct path provided
			allScriptPaths = append(allScriptPaths, scriptField)
		} else {
			// Try to locate the script via `find`
			fmt.Println(scriptField)
			findCmd := exec.Command("sh", "-c", fmt.Sprintf("find / -name '%s' 2>/dev/null", scriptField))
			foundPath, err := findCmd.Output()

			if err != nil && len(foundPath) == 0 {
				gologger.Error().Msgf("Failed to locate script [%s]: %v", scriptField, err)
				continue
			}

			scriptPaths := strings.Split(strings.TrimSpace(string(foundPath)), "\n")
			for _, p := range scriptPaths {
				p = strings.TrimSpace(p)
				if p != "" {
					allScriptPaths = append(allScriptPaths, p)
				}
			}
		}

		for _, scriptPath := range allScriptPaths {
			// --- Permission Check ---
			info, statErr := os.Stat(scriptPath)
			if statErr != nil {
				gologger.Error().Msgf("Unable to stat file %s: %s", scriptPath, statErr)
				continue
			}

			mode := info.Mode()
			if mode.Perm()&0002 != 0 {
				gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs overwrite for %s", scriptPath)
			} else {
				gologger.Print().Label(utils.Sad.String()).Msgf("Unable to overwite for %s", scriptPath)
			}
			gologger.Print().Label(utils.Bsh.String()).Msgf("Use `ls -l %s` to confirm world-writable permission", scriptPath)

			// --- Wildcard Check ---
			data, readErr := os.ReadFile(scriptPath)
			if readErr != nil {
				gologger.Error().Msgf("Unable to read cron script [%s]: %s", scriptPath, readErr)
				continue
			}

			if strings.Contains(string(data), "*") {
				gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs wildcard in %s", scriptPath)
			} else {
				gologger.Print().Label(utils.Sad.String()).Msgf("No wildcard found in script %s", scriptPath)
			}
			fmt.Println()
		}
	}
	fmt.Println()
	gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `cat /etc/crontab`")
}
