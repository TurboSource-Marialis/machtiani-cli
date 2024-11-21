package cli

import (
    "os"
)

func parseFlags(fs *flag.FlagSet, args []string) {
    err := utils.ParseFlags(fs, args) // Parse flags after the command
    if err != nil {
        log.Fatalf("Error parsing flags: %v", err)
    }
}

// handleError prints the error message and exits the program.
func handleError(message string) {
    fmt.Fprintf(os.Stderr, "%s\n", message)
    os.Exit(1)
}
