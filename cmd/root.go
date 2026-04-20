package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

// This is the root command for the db-backup tool
var rootCmd = &cobra.Command{
    Use:   "db-backup",
    Short: "A CLI tool to backup and restore databases",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}