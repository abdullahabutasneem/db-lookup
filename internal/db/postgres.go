package db

import (
	"fmt"
	"os/exec"
)

type Postgres struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func (p *Postgres) Dump(outputPath string) error {
	fmt.Println("Running pg_dump...")
	cmd := exec.Command("pg_dump",
		"-h", p.Host,
		"-U", p.User,
		"-d", p.Name,
		"-f", outputPath,
	)
	cmd.Env = append(cmd.Env, "PGPASSWORD="+p.Password)
	return cmd.Run()
}

func (p *Postgres) Restore(inputPath string) error {
	// implement later
	return nil
}
