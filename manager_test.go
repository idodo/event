package event

import (
	"regexp"
	"testing"
)

func TestGoodName(t *testing.T) {
	var goodNameReg = regexp.MustCompile(`^[0-9a-zA-Z][\w-.*]*$`)
	if goodNameReg.MatchString("9_GROUP") {
		t.Log("match")
	} else {
		t.Error("not match")
	}
}
