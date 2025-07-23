package coreLinux

import (
	"XCalate/engine/utils"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func CheckCronJobs() {
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
}
