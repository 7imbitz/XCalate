package coreLinux

import (
	"XCalate/engine/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func CheckCronJobDetails() {
	gologger.Info().Msg("Checking cronjobs environment PATH ↓")
	gologger.Info().Msg("Checking cronjobs being run with a wildcard (*) ↓")
	gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `cat /etc/crontab`")
	fmt.Println()

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

	// ===== PATH Check =====
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

	// ===== Content Display =====
	gologger.Print().Label(utils.Res.String()).Msg("Content of /etc/crontab")
	lines := strings.Split(crontabString, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			continue
		}
		fmt.Println(line)
	}

	//	var customScripts []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Contains(line, "run-parts") {
			continue // skip built-in cron jobs
		}

		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		//scriptField := fields[5]
		scriptField := strings.Join(fields[6:], " ")
		//fmt.Println("scriptField", scriptField)

		// If full path
		if strings.HasPrefix(scriptField, "/") {
			fullPath := scriptField
			//fmt.Println("fullPath", fullPath)
			data, err := os.ReadFile(fullPath)
			if err != nil {
				gologger.Error().Msgf("Unable to read cron script [%s]: %s", fullPath, err)
				continue
			}

			if strings.Contains(string(data), "*") {
				gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs wildcard in %s", fullPath)
			} else {
				gologger.Print().Label(utils.Sad.String()).Msgf("No wildcard found in script %s", fullPath)
			}
		} else {
			findCmd := exec.Command("sh", "-c", fmt.Sprintf("find / -name '%s' 2>/dev/null", scriptField))
			foundPath, err := findCmd.Output()
			if err != nil || len(foundPath) == 0 {
				gologger.Error().Msgf("Failed to locate script [%s]: %s", scriptField, err)
				log.Fatalf("error %v", err)
				continue
			}

			scriptPaths := strings.Split(strings.TrimSpace(string(foundPath)), "\n")
			for _, scriptPath := range scriptPaths {
				scriptPath = strings.TrimSpace(scriptPath)

				info, statErr := os.Stat(scriptPath)
				if statErr != nil {
					gologger.Error().Msgf("Unable to stat file %s: %s", scriptPath, statErr)
					continue
				}

				mode := info.Mode()
				if mode.Perm()&0002 != 0 {
					// Only mark as cron overwrite possibility if world-writable
					gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs overwrite for %s", scriptField)
					gologger.Print().Label(utils.Bsh.String()).Msgf("Use `ls -l %s` to confirm world-writable permission", scriptPath)
				}
			}

			fullPath := strings.TrimSpace(string(foundPath))
			data, err := os.ReadFile(fullPath)
			if err != nil {
				gologger.Error().Msgf("Unable to read cron script [%s]: %s", fullPath, err)
				continue
			}
			if strings.Contains(string(data), "*") {
				gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs wildcard in %s", fullPath)
			} else {
				gologger.Print().Label(utils.Sad.String()).Msgf("No wildcard found in script %s", fullPath)
			}
		}
	}

}

/*
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if !strings.Contains(line, "/") && !strings.HasSuffix(line, ".sh") {
			continue
		}
		if strings.Contains(line, "run-parts") {
			continue // skip built-in run-parts cron jobs
		}

		fields := strings.Fields(line)
		if len(fields) >= 6 {
			scriptField := fields[5]
			//scriptName := path.Base(scriptField)
			scriptName := scriptField // preserve full path
			customScripts = append(customScripts, scriptName)
		}
	}

	if len(customScripts) == 0 {
		gologger.Print().Label(utils.Sad.String()).Msg("No custom cron jobs found in /etc/crontab")
		return
	}

	for _, scriptName := range customScripts {
		var fullPath string

		// If script already has a full path, use it directly
		if strings.HasPrefix(scriptName, "/") {
			fullPath = scriptName
		} else {
			// Try to locate the script using `find`
			findCmd := exec.Command("sh", "-c", fmt.Sprintf("find / -name '%s' 2>/dev/null", scriptName))
			foundPath, err := findCmd.Output()
			if err != nil || len(foundPath) == 0 {
				if strings.Contains(pathLine, currentUser.HomeDir) {
					gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs overwrite for %s", scriptName)
				} else {
					gologger.Error().Msgf("Failed to locate script [%s]: %s", scriptName, err)
				}
				continue
			}
			fullPath = strings.TrimSpace(string(foundPath))
		}

		// Read the script content
		fmt.Println("I'm here!")
		fmt.Println(fullPath)
		data, err := os.ReadFile(fullPath)
		if err != nil {
			gologger.Error().Msgf("Unable to read cron script [%s]: %s", fullPath, err)
			continue
		}

		if strings.Contains(string(data), "*") {
			gologger.Print().Label(utils.Res.String()).Msgf("Possible cron jobs wildcard in %s", fullPath)
		} else {
			gologger.Print().Label(utils.Sad.String()).Msgf("No wildcard found in script %s", fullPath)
		}
	}
*/
/*func CheckCronJobs() {
	content, err := os.ReadFile(utils.CronPath)
	if err != nil {
		gologger.Error().Msgf("Error reading /etc/crontab: %s", err)
		return
	}

	gologger.Print().Label(utils.Res.String()).Msg("Content of /etc/crontab")
	// Convert content to string
	crontabContent := string(content)

	// Split the content into lines
	lines := strings.Split(crontabContent, "\n")

	// Print non-comment lines
	for _, line := range lines {
		// Skip lines starting with #
		if strings.HasPrefix(line, "#") {
			continue
		}
		fmt.Println(line)
	}
}

// Check cronjobs PATH if it contain $HOME
func CheckCronjobsPath() {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		gologger.Error().Msgf("Error: %s", err)
		return
	}

	// Read the contents of /etc/crontab
	crontabContent, err := os.ReadFile(utils.CronPath)
	if err != nil {
		gologger.Error().Msgf("Error reading /etc/crontab: %s", err)
		return
	}

	// Convert the content to string
	crontabString := string(crontabContent)

	// Find the line with "PATH=" in /etc/crontab
	pathLine := ""
	for _, line := range strings.Split(crontabString, "\n") {
		if strings.HasPrefix(line, "PATH=") {
			pathLine = line
			break
		}
	}

	// Check if the home directory is in the PATH
	if strings.Contains(pathLine, currentUser.HomeDir) {
		gologger.Print().Label(utils.Res.String()).Msg("PATH in /etc/crontab has home user directory")
		gologger.Print().Label(utils.Bsh.String()).Msg("Check out this command to verify! `cat /etc/crontab`")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg("PATH in /etc/crontab does not have home user directory")
	}
}*/
