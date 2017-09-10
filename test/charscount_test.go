package test

import "testing"
import "../milestones"

func TestCharscountWithAnd(t *testing.T) {
	outWithAnd := 21124
	if x := milestones.CharsCount(true); x != outWithAnd {
		t.Errorf("Lettercount returned %d but we should have %d", x, outWithAnd)
	}
}

func TestCharsountWithoutAnd(t *testing.T) {
	outWithoutAnd := 18451
	if x := milestones.CharsCount(false); x != outWithoutAnd {
		t.Errorf("Lettercount returned %d but we should have %d", x, outWithoutAnd)
	}

}