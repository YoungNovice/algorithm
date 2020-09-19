package brackets

import (
	"testing"
)

func TestMatch(t *testing.T) {
	b := MatchComplex("(asdf(s)(())()){(){}[nihao]}")
	if !b {
		t.Error("not match")
	} else {
		t.Log("match")
	}
}