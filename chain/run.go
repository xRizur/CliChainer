package main

import (
    "github.com/spf13/cobra"
    "log"
    "strings"
	"fmt"
)


var runCmd = &cobra.Command{
    Use:   "run [alias] [args...]",
    Short: "Executes a command chain defined for an alias",
    Long: `Executes a series of commands defined for the given alias in config.yaml.
           Additional arguments can be passed which will be used in the commands.`,
    Args:  cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        alias := args[0]
        commandArgs := args[1:]

        config, err := ReadConfig("config.yaml")
        if err != nil {
            log.Fatalf("Błąd odczytu konfiguracji: %v", err)
        }

        aliasConfig, ok := config.Aliases[alias]
        if !ok {
            log.Fatalf("Nie znaleziono aliasu: %s", alias)
        }

        for _, command := range aliasConfig.Commands {
			updatedArgs := replaceDynamicPlaceholders(command.Args, commandArgs, config.Variables)

            if err := executeCommand(Command{Name: command.Name, Args: updatedArgs}); err != nil {
                log.Printf("Błąd podczas wykonywania komendy: %v", err)
            }
        }
    },
}

func replaceDynamicPlaceholders(args []string, commandArgs []string, variables map[string]map[string]string) []string {
    var updatedArgs []string
    for _, arg := range args {
        updatedArg := arg
        for i, commandArg := range commandArgs {
            placeholder := fmt.Sprintf("{arg%d}", i+1)
            updatedArg = strings.ReplaceAll(updatedArg, placeholder, commandArg)

            for varCategory, varMappings := range variables {
                dynamicPlaceholder := fmt.Sprintf("{%s.$arg}", varCategory)
                if value, ok := varMappings[commandArg]; ok {
                    updatedArg = strings.ReplaceAll(updatedArg, dynamicPlaceholder, value)
                }
            }
        }
        updatedArgs = append(updatedArgs, updatedArg)
    }
    return updatedArgs
}
