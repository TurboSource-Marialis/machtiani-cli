// runAicommit generates a commit message using aicommit and lets it perform the git commit.
func commit(args []string) {
    config, err := utils.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Define flags specific to aicommit
    fs := flag.NewFlagSet("aicommit", flag.ExitOnError)
    openaiKey := fs.String("openai-key", config.Environment.ModelAPIKey, "OpenAI API Key")
    modelFlag := fs.String("model", "gpt-4o-mini", "Model to use for generating messages")
    amend := fs.Bool("amend", false, "Amend the last commit instead of creating a new one")
    context := fs.String("context", "", "Additional context for generating the commit message")

    // Parse the provided arguments
    err = fs.Parse(args)
    if err != nil {
        handleError(fmt.Sprintf("Error parsing flags: %v", err))
    }

    // Construct aicommit arguments without --dry
    aicommitArgs := []string{
        "--openai-key", *openaiKey,
        "--model", *modelFlag,
    }
    if *amend {
        aicommitArgs = append(aicommitArgs, "--amend")
    }
    if *context != "" {
        aicommitArgs = append(aicommitArgs, "--context", *context)
    }

    // Handle dry-run mode by adding --dry to aicommit arguments if needed
    if utils.IsDryRunEnabled() {
        aicommitArgs = append(aicommitArgs, "--dry")
    }

    // Locate the aicommit binary
    binaryPath, err := exec.LookPath("aicommit")
    if err != nil {
        handleError(fmt.Sprintf("aicommit binary not found in PATH: %v", err))
    }

    // Create the command to run aicommit
    cmd := exec.Command(binaryPath, aicommitArgs...)

    // Set the working directory to the current directory
    cwd, err := os.Getwd()
    if err != nil {
        handleError(fmt.Sprintf("Failed to get current working directory: %v", err))
    }
    cmd.Dir = cwd

    // Inherit environment variables
    cmd.Env = os.Environ()

    // Attach stdout and stderr to display aicommit output directly
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Execute the aicommit command
    err = cmd.Run()
    if err != nil {
        handleError(fmt.Sprintf("Error running aicommit: %v", err))
    }

    // No need to perform git commit manually; aicommit handles it
}

// handleError prints the error message and exits the program.
func handleError(message string) {
    fmt.Fprintf(os.Stderr, "%s\n", message)
    os.Exit(1)
}

