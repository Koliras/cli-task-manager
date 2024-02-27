package cmd

import (
  "fmt"
  "errors"

  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
  Use:   "list",
  Short: "Show the list of all your incomplete tasks",
  RunE: func(cmd *cobra.Command, args []string) error {
      fmt.Println("Temporary list content")
      if false {
          return errors.New("Some error") 
      }
      return nil
  },
}
