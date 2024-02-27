package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Koliras/cli-task-manager/db"
	"github.com/spf13/cobra"
)

func init() {
  RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
    Use:   "do",
    Short: "Finish the task from the list",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            fmt.Println("Please, enter the key of the task")
            os.Exit(1)
        }
        if len(args) > 1 {
            fmt.Println("Too many arguments")
            os.Exit(1)
        }
        key, err:= strconv.Atoi(args[0])
        if err != nil {
            fmt.Println("Key must be a valid number")
            os.Exit(1)
        }
        err = db.DeleteTask(key)
        if err != nil {
            fmt.Println("Something went wrong:", err.Error())
            os.Exit(1)
        }
        fmt.Println("Deleted the task from the list")
    },
}
