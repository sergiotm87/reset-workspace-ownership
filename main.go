package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func changeOwnership(path string, uid int) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path does not exist: %s", path)
	}

	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chown(name, uid, -1)
	})
}

func run() error {
	path := os.Getenv("INPUT_PATH")
	uidStr := os.Getenv("INPUT_UID")

	if path == "" {
		return fmt.Errorf("path input is required")
	}

	if uidStr == "" {
		return fmt.Errorf("uid input is required")
	}

	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		return fmt.Errorf("invalid uid: %w", err)
	}

	if err := changeOwnership(path, uid); err != nil {
		return err
	}

	fmt.Printf("Successfully changed ownership of %s and its contents to uid %d\n", path, uid)
	return nil
}


func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
