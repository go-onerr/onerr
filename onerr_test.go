package onerr

import (
	"bytes"
	"errors"
	"log"
	"testing"
)

var (
	errTest    = errors.New("test error")
	nilErrFunc = func() error {
		return nil
	}
	errFunc = func() error {
		return errTest
	}
)

func TestPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r != errTest {
			t.Fatalf("Panic should panic the given error, got: %#v, want: %#v",
				r, errTest)
		}
	}()

	Panic(errTest)
	t.Fatal("Panic(err) should panic")
}

func TestPanicf(t *testing.T) {
	const want = "error 5: test error"
	defer func() {
		r := recover()
		if r != want {
			t.Fatalf("invalid output for Panicf, got: %v, want: %s", r, want)
		}
	}()

	Panicf(errTest, "error %d", 5)
	t.Fatal("Panicf(err) should panic")
}

func TestNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Panic should not panic, got: %#v", r)
		}
	}()

	Panic(nil)
	Panicf(nil, "")
}

func TestDefaultLogger(t *testing.T) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetFlags(0)

	Log(errTest)

	got := buf.String()
	want := errTest.Error() + "\n"
	if got != want {
		t.Errorf("Log.Print printed %q, want %q", got, want)
	}
}

func TestLog(t *testing.T) {
	buf := setTestLogger()
	defer restoreLogger()

	Log(nil)
	if buf.Len() != 0 {
		t.Fatalf("Log(nil) should not output anything, got: %s", buf.String())
	}

	Log(errTest)
	got := buf.String()
	want := errTest.Error()
	if got != want {
		t.Errorf("Log() = %q, want %q", got, want)
	}
}

func TestLogf(t *testing.T) {
	buf := setTestLogger()
	defer restoreLogger()

	Logf(nil, "")
	if buf.Len() != 0 {
		t.Fatalf("Logf(nil) should not output anything, got: %s", buf.String())
	}

	Logf(errTest, "error")
	got := buf.String()
	want := "error: " + errTest.Error()
	if got != want {
		t.Errorf("Logf() = %q, want %q", got, want)
	}
}

func TestLogFunc(t *testing.T) {
	buf := setTestLogger()
	defer restoreLogger()

	LogFunc(nilErrFunc)
	got := buf.String()
	if got != "" {
		t.Fatalf("LogFunc(nil) should not output anything, got: %s", got)
	}

	LogFunc(errFunc)
	got = buf.String()
	want := errTest.Error()
	if got != want {
		t.Errorf("LogFunc() = %q, want %q", got, want)
	}
}

func TestLogFuncf(t *testing.T) {
	buf := setTestLogger()
	defer restoreLogger()

	LogFuncf(nilErrFunc, "")
	if buf.Len() != 0 {
		t.Fatalf("LogFuncf(nil) should not output anything, got: %s", buf.String())
	}

	LogFuncf(errFunc, "error")
	got := buf.String()
	want := "error: " + errTest.Error()
	if got != want {
		t.Errorf("LogFuncf() = %q, want %q", got, want)
	}
}

var defaultLogger = Logger

func setTestLogger() *bytes.Buffer {
	buf := new(bytes.Buffer)
	Logger = func(msg string) {
		buf.WriteString(msg)
	}
	return buf
}

func restoreLogger() {
	Logger = defaultLogger
}
