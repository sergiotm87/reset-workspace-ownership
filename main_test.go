package main

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestChangeOwnership(t *testing.T) {
	// Create a temporary directory
	tmpdir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpdir)

	// Create a subdirectory and a file
	err = os.Mkdir(filepath.Join(tmpdir, "subdir"), 0755)
	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Create(filepath.Join(tmpdir, "subdir", "testfile"))
	if err != nil {
		t.Fatal(err)
	}

	// Get the current user's uid
	uid := os.Getuid()

	// Change ownership
	err = changeOwnership(tmpdir, uid)
	if err != nil {
		t.Fatal(err)
	}

	// Check if the ownership was changed for all files and directories
	err = filepath.Walk(tmpdir, func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		stat, ok := info.Sys().(*syscall.Stat_t)
		if !ok {
			t.Fatal("could not get file stat")
		}

		if int(stat.Uid) != uid {
			t.Errorf("expected uid %d, got %d for %s", uid, stat.Uid, name)
		}
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestInvalidPath(t *testing.T) {
	err := changeOwnership("/non/existent/path", 1000)
	if err == nil {
		t.Fatal("expected error, but got nil")
	}

	expected := "path does not exist: /non/existent/path"
	if err.Error() != expected {
		t.Errorf("expected error %q, got %q", expected, err.Error())
	}
}
