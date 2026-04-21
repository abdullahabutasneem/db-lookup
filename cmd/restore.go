package cmd

import (
	"fmt"

	"github.com/abdullahabutasneem/db-backup/internal/db"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore a database from a SQL file",
	Run:   runRestore,
}

var restoreHost string
var restoreName string
var restoreUser string
var restorePassword string
var restoreFile string
var psqlPath string

func init() {
	restoreCmd.Flags().StringVar(&restoreHost, "host", "localhost", "Database host")
	restoreCmd.Flags().StringVar(&restoreName, "name", "", "Database name")
	restoreCmd.Flags().StringVar(&restoreUser, "user", "postgres", "Database user")
	restoreCmd.Flags().StringVar(&restorePassword, "password", "secret", "Database password")
	restoreCmd.Flags().StringVar(&restoreFile, "file", "./backup.sql", "Path to SQL backup file")
	restoreCmd.Flags().StringVar(&psqlPath, "psql-path", "", "Path to psql binary (optional)")
	restoreCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(restoreCmd)
}

func runRestore(cmd *cobra.Command, args []string) {
	pg := &db.Postgres{
		Host:     restoreHost,
		Name:     restoreName,
		User:     restoreUser,
		Password: restorePassword,
		PsqlPath: psqlPath,
	}

	err := pg.Restore(restoreFile)
	if err != nil {
		fmt.Println("Restore failed:", err)
		return
	}
	fmt.Printf("Restore complete from: %s\n", restoreFile)
}
