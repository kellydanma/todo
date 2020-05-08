package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	// build binary
	fmt.Println("Building binary...")
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}
	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Could not build binary %s: %s", binName, err)
		os.Exit(1)
	}

	// run tests
	fmt.Println("Running tests...")
	result := m.Run()

	// clean up
	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)
	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "clean the kitchen"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	cmdPath := filepath.Join(dir, binName)

	// test adding tasks
	t.Run("Add new task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, strings.Split(task, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	// test listing tasks
	t.Run("List tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := fmt.Sprintf("1. %s\n", task)
		if expected != string(out) {
			t.Errorf("Expected %q, got %q instead\n", expected, string(out))
		}
	})
}
