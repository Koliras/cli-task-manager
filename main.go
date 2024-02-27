package main

import (
    "fmt"
    "os"
    "path/filepath"

    homedir "github.com/mitchellh/go-homedir"
    "github.com/Koliras/cli-task-manager/cmd"
    "github.com/Koliras/cli-task-manager/db"
)

func main() {
    home, err := homedir.Dir()
    if err != nil {
        fmt.Println("Could not find Go home directory")
        os.Exit(1)
    }
    dbPath := filepath.Join(home, "tasks.db")
    err = db.Init(dbPath)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    cmd.RootCmd.Execute()
}
