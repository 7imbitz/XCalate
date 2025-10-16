package coreLinux

import (
	"XCalate/engine/report"
	"XCalate/engine/utils"
	"os"
	"os/user"
	"path/filepath"

	"github.com/projectdiscovery/gologger"
)

func CheckFilesAndDirs() {
	currentUser, err := user.Current()
	if err != nil {
		gologger.Error().Msgf("Error getting current user: %v\n", err)
		return
	}

	homeDir := currentUser.HomeDir

	// Check for any *history file in home dir
	matches, err := filepath.Glob(filepath.Join(homeDir, ".*history"))
	if err != nil {
		gologger.Error().Msgf("Error during globbing history files: %v\n", err)
	} else if len(matches) > 0 {
		report.MarkTask("Task 16 - Passwords & Keys - History Files", true)
		for _, file := range matches {
			gologger.Print().Label(utils.Res.String()).Msgf("history file exists: %s\n", file)
		}
	} else {
		report.MarkTask("Task 16 - Passwords & Keys - History Files", false)
		gologger.Print().Label(utils.Sad.String()).Msg("no history file found in home directory")
	}

	// Check for .ssh in home directory
	sshHome := filepath.Join(homeDir, ".ssh")
	if stat, err := os.Stat(sshHome); err == nil && stat.IsDir() {
		gologger.Print().Label(utils.Res.String()).Msg(".ssh directory exists in ~")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg(".ssh directory NOT found in ~")
	}

	// Check for .ssh in root (/)
	sshRoot := "/.ssh"
	if stat, err := os.Stat(sshRoot); err == nil && stat.IsDir() {
		gologger.Print().Label(utils.Res.String()).Msg(".ssh directory exists in /")
	} else {
		gologger.Print().Label(utils.Sad.String()).Msg(".ssh directory NOT found in /")
	}
	report.MarkTaskStatus("Task 17 - Passwords & Keys - Config Files", report.Manual)
	report.MarkTaskStatus("Task 18 - Passwords & Keys - SSH Keys", report.Manual)
}
