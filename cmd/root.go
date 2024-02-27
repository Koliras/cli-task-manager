package cmd

import (
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "cli-task-manager",
    Short: "Task is a CLI task manager",
}

