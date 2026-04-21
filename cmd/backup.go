package cmd

import (
	"fmt"

	"github.com/abdullahabutasneem/db-backup/internal/db"
	"github.com/spf13/cobra"
)

var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup a database",
	Run:   runBackup, // <-- this is the handler for the backup command
}

// These are like query params / request body
var dbHost string
var dbName string
var dbUser string
var dbPassword string
var pgDumpPath string

func init() {
	// Register flags — like defining your API schema
	backupCmd.Flags().StringVar(&dbHost, "host", "localhost", "Database host")
	backupCmd.Flags().StringVar(&dbName, "name", "", "Database name")
	backupCmd.Flags().StringVar(&dbUser, "user", "postgres", "Database user")
	backupCmd.Flags().StringVar(&dbPassword, "password", "secret", "Database password")
	backupCmd.Flags().StringVar(&pgDumpPath, "pg-dump-path", "", "Path to pg_dump binary (optional)")
	backupCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(backupCmd) // Register with root — like router.GET("/backup", handler)
}

func runBackup(cmd *cobra.Command, args []string) {
	pg := &db.Postgres{
		Host:     dbHost,
		Name:     dbName,
		User:     dbUser,
		Password: dbPassword,
		DumpPath: pgDumpPath,
	}

	err := pg.Dump("./backup.sql")
	if err != nil {
		fmt.Println("Backup failed:", err)
		return
	}
	fmt.Println("Backup complete: ./backup.sql")
}
