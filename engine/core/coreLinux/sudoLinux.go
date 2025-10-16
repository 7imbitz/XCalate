package coreLinux

import (
	"XCalate/engine/report"
	"XCalate/engine/utils"
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strings"

	"github.com/projectdiscovery/gologger"
)

func CheckSudoCommands() {
	currentUser, err := user.Current()
	if err != nil {
		gologger.Fatal().Msgf("Failed to get current user: %v", err)
	}

	cmd := exec.Command("sudo", "-l")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		gologger.Error().Msgf(`Failed to run "sudo -l": %v`, err)
		gologger.Print().Label(utils.Sad.String()).Msg(`Skipping sudo privilege checks — continuing with other modules ↓`)
		report.MarkTask("Task 6 - Sudo - Shell Escape Sequences", false)
		return
	}

	output := out.String()
	if strings.TrimSpace(output) == "" {
		gologger.Print().Label(utils.Sad.String()).Msg(`"sudo -l" returned empty output — skipping further analysis ↓`)
		report.MarkTask("Task 6 - Sudo - Shell Escape Sequences", false)
		return
	}

	gologger.Info().Msg("Parsing sudo -l output...")
	report.MarkTask("Task 6 - Sudo - Shell Escape Sequences", true)
	parseSudoL(output, currentUser.Username)
	report.MarkTaskStatus("Task 7 - Sudo - Environment Variables", report.Manual)
}

func parseSudoL(output, username string) {
	scanner := bufio.NewScanner(strings.NewReader(output))

	var allPattern = "(ALL:ALL)ALL"

	startParsing := false

	// match explicit NOPASSWD: entries like "(user : group) NOPASSWD: /usr/bin/XXX, /usr/bin/YYY"
	nopassRegex := regexp.MustCompile(`^\(([^)]+)\)\s+NOPASSWD:\s*(.+)`)

	// generic matcher for lines that start with "(...)" followed by a command (captures the entire RHS)
	cmdRegex := regexp.MustCompile(`^\(([^)]+)\)\s*(.+)`)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "User "+username+" may run the following commands") {
			startParsing = true
			continue
		}

		if startParsing {
			line = strings.TrimSpace(line)
			if line == "" {
				break // End of list
			}

			// Normalize spaces for easy ALL-detection:
			normalized := strings.ReplaceAll(line, " ", "")

			// Detect "(ALL : ALL) ALL" and variants like "(ALL) ALL", possibly with NOPASSWD hints
			if strings.Contains(normalized, allPattern) || strings.Contains(normalized, "(ALL)ALL") || strings.Contains(normalized, "NOPASSWD:ALL") {
				gologger.Print().Label(utils.Res.String()).Msgf("User %s able to run ALL", username)
				break
			}

			// First try NOPASSWD style parsing (and split comma-separated commands)
			if matches := nopassRegex.FindStringSubmatch(line); len(matches) == 3 {
				// matches[1] = who, matches[2] = comma separated command list
				binaries := strings.Split(matches[2], ",")
				for _, bin := range binaries {
					gologger.Print().Label(utils.Res.String()).Msgf("User %s may run the %s binary using sudo", username, strings.TrimSpace(bin))
				}
				continue
			}

			// Generic command line: capture whatever is after the "(who)" and print it as a whole
			if matches := cmdRegex.FindStringSubmatch(line); len(matches) == 3 {
				commandStr := strings.TrimSpace(matches[2])
				// If the RHS is a comma-separated list, still print each; otherwise print whole command
				if strings.Contains(commandStr, ",") {
					parts := strings.Split(commandStr, ",")
					for _, p := range parts {
						gologger.Print().Label(utils.Res.String()).Msgf("User %s may run: %s", username, strings.TrimSpace(p))
					}
				} else {
					gologger.Print().Label(utils.Res.String()).Msgf("User %s may run: %s", username, commandStr)
				}
				continue
			}

			// If it reaches here, we couldn't parse the line — print it raw for visibility
			gologger.Print().Label(utils.Sad.String()).Msgf("Unparsed sudo line: %s", line)
		}
	}

	if err := scanner.Err(); err != nil {
		gologger.Fatal().Msgf("Scanner error: %v", err)
	}
}
