package coreLinux

import (
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
		gologger.Fatal().Msgf("Failed to run sudo -l: %v", err)
	}

	output := out.String()
	gologger.Info().Msg("Parsing sudo -l output...")

	parseSudoL(output, currentUser.Username)
}

func parseSudoL(output, username string) {
	scanner := bufio.NewScanner(strings.NewReader(output))

	startParsing := false
	commandRegex := regexp.MustCompile(`\(([^)]+)\)\s+NOPASSWD:\s*(.+)`)

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

			// Match commands
			matches := commandRegex.FindStringSubmatch(line)
			if len(matches) == 3 {
				binaries := strings.Split(matches[2], ",")
				for _, bin := range binaries {
					gologger.Print().Label(utils.Bsh.String()).Msgf("%s\n", strings.TrimSpace(bin))
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		gologger.Fatal().Msgf("Scanner error: %v", err)
	}
}
