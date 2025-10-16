package core

import (
	"XCalate/engine/core/coreLinux"
	"XCalate/engine/report"
	"XCalate/engine/utils"
	"fmt"
	"os/exec"
	"os/user"
	"strings"
	"time"

	"github.com/projectdiscovery/gologger"
)

type Task struct {
	Title   string
	Message string
	Command string
	Exec    func()
}

func LinPrivEscChecker() {
	verifySudo()
	report.PrintSummary()
	gologger.Print().Label(utils.Res.String()).Msgf("Happy hacking!")
	fmt.Println()
}

func profiler() {
	cmds := []struct {
		label string
		name  string
		args  []string
	}{
		{"User", "whoami", nil},
		{"Print Working Dir", "pwd", nil},
		{"User IDs", "id", nil},
		{"Hostname", "hostname", nil},
	}

	for _, c := range cmds {
		out, err := exec.Command(c.name, c.args...).CombinedOutput()
		output := strings.TrimSpace(string(out))
		if err != nil {
			gologger.Error().Msgf("%-12s: ERROR: %v\n", c.label, err)
			if output != "" {
				gologger.Info().Msgf("  output: %s\n", output)
			}
			continue
		}
		gologger.Print().Label(utils.Bsh.String()).Msgf("%-12s: %s\n", c.label, output)
	}
}

func verifySudo() {
	var hasPassword string
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		gologger.Error().Msgf("Error: %s", err)
		return
	}

	gologger.Info().Msg("Starting...")
	profiler()
	gologger.Info().Msg("Checking Sudo user availability ↓")
	gologger.Print().Label(utils.Bsh.String()).Msgf("Do you have %s's password? ((y)es/(n)o) ", currentUser.Username)

	// Read user input
	_, err = fmt.Scanln(&hasPassword)
	if err != nil {
		gologger.Fatal().Msgf("Error reading input: %s", err)
	}

	// Check user input
	switch hasPassword {
	case "yes", "y", "Y":
		gologger.Print().Label(utils.Res.String()).Msg("Running all functions ↓")
		hasSudo()
	case "no", "n", "N":
		gologger.Print().Label(utils.Sad.String()).Msg("Skipping several functions ↓")
		noSudo()
	default:
		gologger.Fatal().Msg("Invalid input. Please enter '(y)es' or '(n)o'.")
	}
}

func runTask(task Task) {
	time.Sleep(3 * time.Second)
	fmt.Println("\n=============== " + task.Title + " ===============")
	gologger.Info().Msg(task.Message)

	if task.Exec != nil {
		task.Exec()
	}

	if task.Command != "" {
		gologger.Print().Msg("Check out this command to verify! `" + task.Command + "`")
	}
}

func hasSudo() {
	tasks := []Task{
		{"Sudo - Shell Escape", "Running sudo -l, you may be prompted for a password... ↓", "", coreLinux.CheckSudoCommands},
		{"Process Exploit", "Checking for process running as root ↓", "", coreLinux.CheckRootProcesses},
		{"/etc/shadow Details", "Checking Shadow file ↓", "", coreLinux.CheckShadowPermissions},
		{"Writable /etc/passwd", "Checkin Passwd file ↓", "", coreLinux.CheckPasswdWritable},
		{"Cron Jobs Details", "Checking cronjobs ↓", "", coreLinux.CheckCronJobDetails},
		{"SUID/SGID Known Exploits", "Listing SUID/SGID Executable ↓", "", coreLinux.CheckSUIDExec},
		{"`Capability` Binary Exploits", "Listing cap_setuid+ep Executable ↓", "", coreLinux.CheckCapability},
		{"Network File System", "Checking for NFS share configuration ↓", "", coreLinux.CheckNFS},
		{"Passwords & Keys", "Checking for password and key files ↓", "", coreLinux.CheckFilesAndDirs},
		{"Kernel Exploit", "Checking for kernel exploit ↓", "", coreLinux.CheckKernel},
	}

	for _, task := range tasks {
		runTask(task)
	}
}

func noSudo() {
	tasks := []Task{
		{"Process Exploit", "Checking for process running as root ↓", "", coreLinux.CheckRootProcesses},
		{"/etc/shadow Details", "Checking Shadow file ↓", "", coreLinux.CheckShadowPermissions},
		{"Writable /etc/passwd", "Checking Passwd file ↓", "", coreLinux.CheckPasswdWritable},
		{"Cron Jobs Details", "Checking cronjobs ↓", "", coreLinux.CheckCronJobDetails},
		{"SUID/SGID Known Exploits", "Listing SUID/SGID Executable ↓", "", coreLinux.CheckSUIDExec},
		{"`Capability` Binary Exploits", "Listing cap_setuid+ep Executable ↓", "", coreLinux.CheckCapability},
		{"Network File System", "Checking for NFS share configuration ↓", "", coreLinux.CheckNFS},
		{"Passwords & Keys", "Checking for password and key files ↓", "", coreLinux.CheckFilesAndDirs},
		{"Kernel Exploit", "Checking for kernel exploit ↓", "", coreLinux.CheckKernel},
	}

	for _, task := range tasks {
		runTask(task)
	}
}
