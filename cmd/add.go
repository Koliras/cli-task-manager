package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Koliras/cli-task-manager/db"
	"github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add the task to the list",
  Run: func(cmd *cobra.Command, args []string) {
      task := strings.Join(args, " ")
      _, err := db.CreateTask(task)
      if err != nil {
          fmt.Println("Something went wrong:", err.Error())
          os.Exit(1)
      }
      fmt.Println("Added task to the list")
  },
}
