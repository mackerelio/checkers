// Package checkers is utility for checker commands
package checkers

import (
	"fmt"
	"os"
)

// Status declare the result of the checker command
type Status int

// Statuses
const (
	OK Status = iota
	WARNING
	CRITICAL
	UNKNOWN
)

func (st Status) String() string {
	switch st {
	case OK:
		return "OK"
	case WARNING:
		return "WARNING"
	case CRITICAL:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// NewChecker returns new Checker
func NewChecker(st Status, msg string) *Checker {
	return &Checker{
		Status:  st,
		Message: msg,
	}
}

// Checker is utility struct for check script
type Checker struct {
	Name    string
	Status  Status
	Message string
}

var exit = os.Exit

// Exit with message
func (ckr *Checker) Exit() {
	fmt.Println(ckr.String())
	exit(int(ckr.Status))
}

func (ckr *Checker) ExitStatusAs(maps map[Status]Status) {
	if _, ok := maps[ckr.Status]; ok {
		ckr.Status = maps[ckr.Status]
	}

	fmt.Println(ckr.String())
	exit(int(ckr.Status))
}

func (ckr *Checker) String() string {
	return fmt.Sprintf("%s %s: %s", ckr.Name, ckr.Status, ckr.Message)
}

// Ok exiting with OK status
func Ok(msg string) *Checker {
	return NewChecker(OK, msg)
}

// Warning exiting with WARNING status
func Warning(msg string) *Checker {
	return NewChecker(WARNING, msg)
}

// Critical exiting with CRITICAL status
func Critical(msg string) *Checker {
	return NewChecker(CRITICAL, msg)
}

// Unknown exiting with UNKNOWN status
func Unknown(msg string) *Checker {
	return NewChecker(UNKNOWN, msg)
}
