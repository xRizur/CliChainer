package main

import (
    "os"
    "os/exec"
)

func executeCommand(cmd Command) error {
    command := exec.Command(cmd.Name, cmd.Args...)
    command.Env = os.Environ()
    command.Stdout = os.Stdout
    command.Stderr = os.Stderr
    return command.Run()
}

