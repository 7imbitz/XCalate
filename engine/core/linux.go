package core

import (
	"XCalate/engine/core/coreLinux"
	"XCalate/engine/utils"
	"fmt"
	"os/user"
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
}

func verifySudo() {
	var hasPassword string
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		gologger.Error().Msgf("Error: %s", err)
		return
	}

	fmt.Println()
	gologger.Info().Msg("Checking Sudo user availability ↓")
	fmt.Printf("Do you have %s's password? (yes/no) ", currentUser.Username)

	// Read user input
	_, err = fmt.Scanln(&hasPassword)
	if err != nil {
		gologger.Fatal().Msgf("Error reading input: %s", err)
	}

	// Check user input
	switch hasPassword {
	case "yes":
		gologger.Print().Label(utils.Res.String()).Msg("Running all functions ↓")
		hasSudo()
	case "no":
		gologger.Print().Label(utils.Sad.String()).Msg("Skipping several functions ↓")
		noSudo()
	default:
		gologger.Fatal().Msg("Invalid input. Please enter 'yes' or 'no'.")
	}
}

func runTask(task Task) {
	time.Sleep(2 * time.Second)
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
		{"Process Exploit", "Checking for process running as root ↓", "", coreLinux.CheckRootProcesses},
		{"Readable /etc/shadow", "Checking if Shadow file is world-readable ↓", "", coreLinux.CheckShadowReadable},
		{"Writable /etc/shadow", "Checking if Shadow file is world-writable ↓", "", coreLinux.CheckShadowWritable},
		{"Writable /etc/passwd", "Checking if Passwd file is world-writable ↓", "", coreLinux.CheckPasswdWritable},
		//{"Sudo - Shell Escape", "Running sudo -l ↓", "", checkSudoCommands},
		//{"Sudo - Environment Variable", "Listing environment variable and shell escape ↓", "", checkShellEscapePath},
		{"Cron Jobs PATH", "Checking cronjobs environment PATH ↓", "", coreLinux.CheckCronjobsPath},
		{"Cron Jobs File", "Checking cronjobs ↓", "", coreLinux.CheckCronJobs},
		//{"Cron Jobs Wildcard", "Need to manually check the content of cron jobs script (if any) ↓", "cat /etc/crontab", nil},
		//{"SUID/SGID Known Exploits", "Listing SUID/SGID Executable ↓", "", checkSuidExec},
		{"SUID/SGID Shared Object Injection", "Verify vulnerable executables ↓", "find / -type f -a \\( -perm -u+s -o -perm -g+s \\) -exec ls -l {} \\; 2> /dev/null", nil},
		{"Passwords & Keys - History Files", "Need to manually verify password or key files ↓", "cat ~history | less", nil},
		{"Passwords & Keys - Config Files", "Find config/password files ↓", "find /etc -name \"*.conf\" -o -name \"*.cfg\" -o -name \"*.config\"", nil},
		{"Passwords & Keys - SSH Keys", "Find SSH private key ↓", "ls -la / OR ls -la /.ssh OR ls -la ~/.ssh", nil},
		//{"Network File System", "Checking for NFS share configuration ↓", "", checkNFS},
		//{"Kernel Exploits", "Checking for possible kernel exploit ↓", "", checkKernel},
	}

	for _, task := range tasks {
		runTask(task)
	}
}

func noSudo() {
	tasks := []Task{
		{"Process Exploit", "Checking for process running as root ↓", "", coreLinux.CheckRootProcesses},
		{"Readable /etc/shadow", "Checking if Shadow file is world-readable ↓", "", coreLinux.CheckShadowReadable},
		{"Writable /etc/shadow", "Checking if Shadow file is world-writable ↓", "", coreLinux.CheckShadowWritable},
		{"Writable /etc/passwd", "Checking if Passwd file is world-writable ↓", "", coreLinux.CheckPasswdWritable},
		{"Cron Jobs PATH", "Checking cronjobs environment PATH ↓", "", coreLinux.CheckCronjobsPath},
		{"Cron Jobs File", "Checking cronjobs ↓", "", coreLinux.CheckCronJobs},
		//{"Cron Jobs Wildcard", "Need to manually check cron jobs script ↓", "cat /etc/crontab", nil},
		//{"SUID/SGID Known Exploits", "Listing SUID/SGID Executable ↓", "", checkSuidExec},
		{"Passwords & Keys - History Files", "Verify password/key files ↓", "cat ~history | less", nil},
		{"Passwords & Keys - Config Files", "Find config/password files ↓", "find /etc -name \"*.conf\" -o -name \"*.cfg\" -o -name \"*.config\"", nil},
		{"Passwords & Keys - SSH Keys", "Find SSH private key ↓", "ls -la / OR ls -la /.ssh OR ls -la ~/.ssh", nil},
		//{"Network File System", "Checking for NFS share configuration ↓", "", checkNFS},
		//{"Kernel Exploits", "Checking for possible kernel exploit ↓", "", checkKernel},
	}

	for _, task := range tasks {
		runTask(task)
	}
}
