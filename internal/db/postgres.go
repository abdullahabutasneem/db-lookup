package db

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
)

type Postgres struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	DumpPath string
}

func (p *Postgres) Dump(outputPath string) error {
	pgDumpPath, err := resolvePgDumpPath(p.DumpPath)
	if err != nil {
		return err
	}

	fmt.Println("Running pg_dump...")
	cmd := exec.Command(pgDumpPath,
		"-h", p.Host,
		"-U", p.User,
		"-d", p.Name,
		"-f", outputPath,
	)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+p.Password)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (p *Postgres) Restore(inputPath string) error {
	// implement later
	return nil
}

func resolvePgDumpPath(customPath string) (string, error) {
	if customPath != "" {
		if _, err := os.Stat(customPath); err != nil {
			return "", fmt.Errorf("pg_dump binary not found at %q: %w", customPath, err)
		}
		return customPath, nil
	}

	if path, err := exec.LookPath("pg_dump"); err == nil {
		return path, nil
	}

	if runtime.GOOS == "windows" {
		candidates := make([]string, 0)
		programFiles := []string{
			`C:\Program Files\PostgreSQL\*\bin\pg_dump.exe`,
			`C:\Program Files (x86)\PostgreSQL\*\bin\pg_dump.exe`,
		}

		for _, pattern := range programFiles {
			matches, _ := filepath.Glob(pattern)
			candidates = append(candidates, matches...)
		}

		if len(candidates) > 0 {
			sort.Strings(candidates)
			return candidates[len(candidates)-1], nil
		}
	}

	return "", fmt.Errorf("pg_dump not found. Install PostgreSQL client tools and add pg_dump to PATH, or pass --pg-dump-path")
}
