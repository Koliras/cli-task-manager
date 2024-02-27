package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Koliras/cli-task-manager/db"
	"github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(editCmd)
}

var editCmd = &cobra.Command{
  Use:   "edit",
  Short: "Edit the task content",
  Run: func(cmd *cobra.Command, args []string) {
      if len(args) < 2 {
          fmt.Println("There should be at least 2 arguments: 1 for key and 1 for content")
          os.Exit(1)
      }
      key, err := strconv.Atoi(args[0])
      if err != nil {
          fmt.Println("First argument should be a valid task key")
          os.Exit(1)
      }
      newContent := strings.Join(args[1:], " ")
      err = db.EditTask(key, newContent)
      if err != nil {
          fmt.Println("Something went wrong:", err.Error())
          os.Exit(1)
      }
      fmt.Println("Edited the task of the list")
  },
}
