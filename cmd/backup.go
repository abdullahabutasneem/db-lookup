package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
    Use:   "backup",
    Short: "Backup a database",
    Run:   runBackup,   // <-- this is the handler for the backup command
}

// These are like query params / request body
var dbHost string
var dbName string

func init() {
    // Register flags — like defining your API schema
    backupCmd.Flags().StringVar(&dbHost, "host", "localhost", "Database host")
    backupCmd.Flags().StringVar(&dbName, "name", "", "Database name")
    backupCmd.MarkFlagRequired("name")

    rootCmd.AddCommand(backupCmd) // Register with root — like router.GET("/backup", handler)
}

func runBackup(cmd *cobra.Command, args []string) {
    // For now, just print. You'll replace this with real logic later.
    fmt.Printf("Backing up database '%s' on host '%s'\n", dbName, dbHost)
}