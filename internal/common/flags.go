package common

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Function to get the value of a commandline flag by its name.
// If the flag doesn't exist, returns an empty string.
func GetFlag(flagName string) string {
	argv := os.Args[1:]
	argc := len(argv)

	log.Tracef("Searching for %s in %v", flagName, argv)

	for i := 0; i < argc; i++ {
		arg := argv[i]
		var value string

		// Check if the current argument is a commandline flag.
		if strings.HasPrefix(arg, "-") {
			flagPrefix := "-"
			if arg[1] == '-' {
				flagPrefix = "--"
			}

			if arg == flagPrefix+flagName {
				if argc > i+1 {
					value = argv[i+1]
				}
			} else if strings.HasPrefix(arg, flagPrefix+flagName+"=") {
				value = strings.Split(arg, "=")[1]
			}

			if value != "" {
				log.Tracef("foung flag %s=%s", flagName, value)
				return value
			}
		}
	}

	return ""
}
