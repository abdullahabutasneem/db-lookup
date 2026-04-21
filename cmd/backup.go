package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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
var backupOutputDir string

func init() {
	// Register flags — like defining your API schema
	backupCmd.Flags().StringVar(&dbHost, "host", "localhost", "Database host")
	backupCmd.Flags().StringVar(&dbName, "name", "", "Database name")
	backupCmd.Flags().StringVar(&dbUser, "user", "postgres", "Database user")
	backupCmd.Flags().StringVar(&dbPassword, "password", "secret", "Database password")
	backupCmd.Flags().StringVar(&pgDumpPath, "pg-dump-path", "", "Path to pg_dump binary (optional)")
	backupCmd.Flags().StringVar(&backupOutputDir, "output-dir", "./backups", "Directory where backup SQL files are saved")
	backupCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(backupCmd) // Register with root — like router.GET("/backup", handler)
}

func runBackup(cmd *cobra.Command, args []string) {
	if err := os.MkdirAll(backupOutputDir, 0o755); err != nil {
		fmt.Println("Backup failed:", err)
		return
	}

	timestamp := time.Now().Format("20060102-150405")
	backupFilename := fmt.Sprintf("backup-%s.sql", timestamp)
	backupPath := filepath.Join(backupOutputDir, backupFilename)

	pg := &db.Postgres{
		Host:     dbHost,
		Name:     dbName,
		User:     dbUser,
		Password: dbPassword,
		DumpPath: pgDumpPath,
	}

	err := pg.Dump(backupPath)
	if err != nil {
		fmt.Println("Backup failed:", err)
		return
	}
	fmt.Println("Backup complete:", backupPath)
}
