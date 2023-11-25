package main

import (
    "github.com/spf13/cobra"
    "log"
)

var rootCmd = &cobra.Command{
    Use:   "chain",
    Short: "ChainCLI - a tool for executing command chains",
    Long: `ChainCLI is a tool for executing predefined chains of commands. 
           It allows for automation of a series of related CLI commands across different tools.`,
}

func main() {
    rootCmd.AddCommand(runCmd)

    if err := rootCmd.Execute(); err != nil {
        log.Fatalf("Błąd: %v", err)
    }
}
