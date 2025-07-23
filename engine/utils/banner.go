package utils

import (
	"fmt"

	"github.com/projectdiscovery/gologger"
)

// Purely for banner
const Version = `0.1.0`

const Author = `7imbitz`

var banner = fmt.Sprintf(`
   _  ________      __      __     
  | |/ / ____/___ _/ /___ _/ /____ 
  |   / /   / __ '/ / __ '/ __/ _ \
 /   / /___/ /_/ / / /_/ / /_/  __/
/_/|_\____/\__,_/_/\__,_/\__/\___/ 
                                   							  
                                     %s
`, Author)

func ShowBanner() {
	gologger.Print().Msgf("%s", banner)
	gologger.Info().Msgf("XCalate version %s", Version)
	gologger.Info().Msg("A privilege escalation tools based on Top 10 TryHackMe.")
}
