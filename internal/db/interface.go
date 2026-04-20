package db

type Database interface {
	Dump(outputPath string) error
	Restore(inputPath string) error
}
