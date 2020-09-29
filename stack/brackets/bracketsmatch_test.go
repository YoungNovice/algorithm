package brackets

import (
	"testing"
)

func TestMatch(t *testing.T) {
	b := MatchComplex("(])")
	if !b {
		t.Error("not match")
	} else {
		t.Log("match")
	}
}