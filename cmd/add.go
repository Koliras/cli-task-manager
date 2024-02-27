package cmd

import (
  "fmt"
  "errors"

  "github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(tryCmd)
}

var tryCmd = &cobra.Command{
  Use:   "add",
  Short: "Add the task to the list",
  RunE: func(cmd *cobra.Command, args []string) error {
      fmt.Println("Temporary add content")
      if false {
          return errors.New("Some error") 
      }
      return nil
  },
}
