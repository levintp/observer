// This plugin exports functionality to query information about the processes
// running on a node, similar to the `ps` linux utility.
package ps

import (
	"os"
	"regexp"
	"strconv"

	"github.com/levintp/observer/internal/types"
)

const procDir = "/proc"

type pid int

// This exported plugin function counts processes on the local machine.
//
// # Arguments
//
// filter: only count processes that match this regex filter.
func Count(arguments types.ModuleArguments) ([]types.MetricData, error) {
	data := make([]types.MetricData, 0)

	// Get a list of all running processes.
	processes, err := getProcesses()
	if err != nil {
		return nil, err
	}

	filter, hasFilter := arguments["filter"]
	if hasFilter {
		for i, processId := range processes {
			cmdline, err := processId.cmdline()
			if err != nil {
				return nil, err
			}

			match, err := regexp.Match(filter, cmdline)
			if err != nil {

			}
		}
	}

	return data, nil
}

// This function collects all the processes that are running on the local machine.
func getProcesses() ([]pid, error) {
	entries, err := os.ReadDir(procDir)
	if err != nil {
		return nil, err
	}

	pids := make([]pid, 0)

	// Iterate over all the entries in `/proc`.
	for _, entry := range entries {
		// Check if the current entry is a directory.
		if entry.IsDir() {
			// Check if the name of the current directory is a number.
			if processId, err := strconv.Atoi(entry.Name()); err == nil {
				pids = append(pids, pid(processId))
			}
		}
	}

	return pids, nil
}

func (proc pid) cmdline() ([]byte, error) {

}
