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
func (ch *Checker) Exit() {
	fmt.Printf("%s %s\n", ch.Status.String(), ch.Message)
	os.Exit(int(ch.Status))
}

// Ok exiting with OK status
func Ok(msg string) {
	(&Checker{OK, msg}).Exit()
}

// Warning exiting with WARNING status
func Warning(msg string) {
	(&Checker{WARNING, msg}).Exit()
}

// Critical exiting with CRITICAL status
func Critical(msg string) {
	(&Checker{CRITICAL, msg}).Exit()
}

// Unknown exiting with UNKNOWN status
func Unknown(msg string) {
	(&Checker{UNKNOWN, msg}).Exit()
}
