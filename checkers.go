// checkers is utility for checker commands
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
	switch {
	case st == 0:
		return "OK"
	case st == 1:
		return "WARNING"
	case st == 2:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

// Checker is utility struct for check script
type Checker struct {
	Status  Status
	Message string
}

// Exit with message
func (ckr *Checker) Exit() {
	fmt.Println(ckr.String())
	os.Exit(int(ckr.Status))
}

func (ckr *Checker) String() string {
	return fmt.Sprintf("%s %s", ckr.Status.String(), ckr.Message)
}

func (ckr *Checker) Merge(ckrs ...*Checker) *Checker {
	for _, c := range ckrs {
		if c.Status > ckr.Status {
			ckr.Status = c.Status
		}
		ckr.Message = ckr.Message + " - " + c.Message
	}
	return ckr
}

// Ok exiting with OK status
func Ok(msg string) *Checker {
	return &Checker{OK, msg}
}

// Warning exiting with WARNING status
func Warning(msg string) *Checker {
	return &Checker{WARNING, msg}
}

// Critical exiting with CRITICAL status
func Critical(msg string) *Checker {
	return &Checker{CRITICAL, msg}
}

// Unknown exiting with UNKNOWN status
func Unknown(msg string) *Checker {
	return &Checker{UNKNOWN, msg}
}
