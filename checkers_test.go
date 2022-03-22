package checkers

import "testing"

func TestCheck(t *testing.T) {
	tests := []struct {
		name   string
		check  func(s string) *Checker
		msg    string
		status Status
		s      string
	}{
		{"ok", Ok, "OK!", OK, " OK: OK!"},
		{"warn", Warning, "warn!", WARNING, " WARNING: warn!"},
		{"crit", Critical, "crit!", CRITICAL, " CRITICAL: crit!"},
		{"unknown", Unknown, "unknown!", UNKNOWN, " UNKNOWN: unknown!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.check(tt.msg)
			if c.Status != tt.status {
				t.Errorf("Status should be %v but got %v", c.Status, tt.status)
			}
			if c.Message != tt.msg {
				t.Errorf("Message should be %s but got %s", c.Message, tt.msg)
			}
			if s := c.String(); s != tt.s {
				t.Errorf("String() should be %q but got %q", tt.s, s)
			}
		})
	}
}

func TestName(t *testing.T) {
	ckr := Ok("OK!")
	ckr.Name = "someChecker"
	if ckr.String() != "someChecker OK: OK!" {
		t.Errorf("ckr.Status should be 'someChecker OK OK!' but:%v\n", ckr)
	}
}

func TestExit(t *testing.T) {
	var exitCode int
	oexit := exit
	exit = func(code int) {
		exitCode = code
	}
	t.Cleanup(func() {
		exit = oexit
	})

	c := Critical("crit!")
	c.Exit()
	if exitCode != int(CRITICAL) {
		t.Errorf("Exit should exit with %d but got %d", CRITICAL, exitCode)
	}
}
