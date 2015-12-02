package onerr_test

import (
	"errors"
	"os"

	"github.com/go-onerr/onerr"
)

func Example() {
	f, err := os.Create("output.txt")
	onerr.Panic(err)             // Panics if the file cannot be created.
	defer onerr.LogFunc(f.Close) // Logs any error while closing the file.

	_, err = f.WriteString("output")
	// Logs any write error like: error while writing "output.txt": disk is full
	onerr.Logf(err, "error while writing %q", f.Name())
}

func Example_Panic() {
	var err error
	onerr.Panic(err) // Does not panic.

	err = errors.New("example error")
	onerr.Panic(err) // Panics.
}

func Example_Panicf() {
	file := "output.txt"
	var err error
	// Does not panic.
	onerr.Panicf(err, "error while saving %q", file)

	err = errors.New("disk is full")
	// Panics with message: error while saving "output.txt": disk is full
	onerr.Panicf(err, "error while saving %q", file)
}

func Example_Log() {
	var err error
	onerr.Log(err) // Does not log anything.

	err = errors.New("example error")
	onerr.Log(err) // Logs: example error
}

func Example_Logf() {
	file := "output.txt"
	var err error
	// Does not panic.
	onerr.Panicf(err, "error while saving %q", file)

	err = errors.New("disk is full")
	// Logs message: error while saving "output.txt": disk is full
	onerr.Panicf(err, "error while saving %q", file)
}

func Example_LogFunc() {
	f, _ := os.Create("output.txt")
	// Logs an error only if f.Close() returns a non-nil error.
	onerr.LogFunc(f.Close)
}

func Example_LogFuncf() {
	f, _ := os.Create("output.txt")
	// Logs an error only if f.Close() returns a non-nil error.
	// Log Message: error while closing "output.txt": Invalid argument
	onerr.LogFuncf(f.Close, "error while closing %q", f.Name())
}
