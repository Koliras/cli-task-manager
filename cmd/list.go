package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
  "github.com/Koliras/cli-task-manager/db"
)

func init() {
  RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "Show the list of all your incomplete tasks",
  Run: func(cmd *cobra.Command, args []string) {
      tasks, err := db.GetAllTasks()
      if err != nil {
          fmt.Println("Something went wrong:", err.Error())
          os.Exit(1)
      }
      if len(tasks) == 0 {
          fmt.Println("You have no tasks. Take a rest")
          return
      }
      fmt.Println("You have the following tasks:")
      for i, task := range tasks {
          fmt.Printf("%d. %s\n", i+1, task.Value)
      }
  },
}
