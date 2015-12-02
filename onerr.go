// Package onerr provides functions to quickly handle errors.
package onerr

import (
	"fmt"
	"log"
)

// Logger is the logger used by the logging functions. By default it uses
// log.Print.
//
// It is not safe to update the Logger and to concurrently use a logging
// function.
var Logger = func(msg string) {
	log.Print(msg)
}

// Panic panics if err is non-nil.
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

// Panicf panics if err is non-nil. It prefixes the panic message with the given
// string. Arguments are handled in the manner of fmt.Printf.
func Panicf(err error, format string, v ...interface{}) {
	if err != nil {
		panic(getMessage(format, v) + ": " + err.Error())
	}
}

// Log logs the error if it is non-nil.
func Log(err error) {
	if err != nil {
		Logger(err.Error())
	}
}

// Logf logs the error if it is non-nil. It prefixes the panic message with the
// given string. Arguments are handled in the manner of fmt.Printf.
func Logf(err error, format string, v ...interface{}) {
	if err != nil {
		Logger(getMessage(format, v) + ": " + err.Error())
	}
}

// LogFunc logs the error if f returns a non-nil error.
func LogFunc(f func() error) {
	if err := f(); err != nil {
		Logger(err.Error())
	}
}

// LogFuncf logs the error if f returns a non-nil error. It prefixes the panic
// message with the given string. Arguments are handled in the manner of
// fmt.Printf.
func LogFuncf(f func() error, format string, v ...interface{}) {
	if err := f(); err != nil {
		Logger(getMessage(format, v) + ": " + err.Error())
	}
}

func getMessage(format string, v []interface{}) string {
	if len(v) == 0 {
		return format
	}
	return fmt.Sprintf(format, v...)
}
