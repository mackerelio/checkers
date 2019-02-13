package checkers

import "testing"

func TestOk(t *testing.T) {
	ckr := Ok("OK!")

	if ckr.Status != OK {
		t.Errorf("ckr.Status should be OK but:%s\n", ckr.Status)
	}
	if ckr.Message != "OK!" {
		t.Errorf("ckr.Message should be 'OK!' but:%s\n", ckr.Message)
	}

	if ckr.String() != " OK: OK!" {
		t.Errorf("ckr.Status should be 'OK OK!' but:%s\n", ckr.String())
	}
}

func TestWarning(t *testing.T) {
	ckr := Warning("warn!")

	if ckr.Status != WARNING {
		t.Errorf("ckr.Status should be WARNING but:%s\n", ckr.Status)
	}
	if ckr.Message != "warn!" {
		t.Errorf("ckr.Message should be 'warn!' but:%s\n", ckr.Message)
	}

	if ckr.String() != " WARNING: warn!" {
		t.Errorf("ckr.Status should be 'WARNING warn!' but:%s\n", ckr.String())
	}
}

func TestCritical(t *testing.T) {
	ckr := Critical("crit!")

	if ckr.Status != CRITICAL {
		t.Errorf("ckr.Status should be CRITICAL but:%s\n", ckr.Status)
	}
	if ckr.Message != "crit!" {
		t.Errorf("ckr.Message should be 'crit!' but:%s\n", ckr.Message)
	}

	if ckr.String() != " CRITICAL: crit!" {
		t.Errorf("ckr.Status should be 'CRITICAL crit!' but:%s\n", ckr.String())
	}
}

func TestUnknown(t *testing.T) {
	ckr := Unknown("unknown!")

	if ckr.Status != UNKNOWN {
		t.Errorf("ckr.Status should be UNKNOWN but:%s\n", ckr.Status)
	}
	if ckr.Message != "unknown!" {
		t.Errorf("ckr.Message should be 'unknown!' but:%s\n", ckr.Message)
	}

	if ckr.String() != " UNKNOWN: unknown!" {
		t.Errorf("ckr.Status should be 'UNKNOWN unknown!' but:%s\n", ckr.String())
	}
}

func TestName(t *testing.T) {
	ckr := Ok("OK!")
	ckr.Name = "someChecker"
	if ckr.String() != "someChecker OK: OK!" {
		t.Errorf("ckr.Status should be 'someChecker OK OK!' but:%s\n", ckr.String())
	}
}
