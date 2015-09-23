package checkers

import "testing"

func TestCheckers(t *testing.T) {
	ckr := Ok("OK!")

	if ckr.Status != OK {
		t.Errorf("ckr.Status should be OK but:%s\n", ckr.Status)
	}

	if ckr.Message != "OK!" {
		t.Errorf("ckr.Message should be 'OK!' but:%s\n", ckr.Message)
	}
}
