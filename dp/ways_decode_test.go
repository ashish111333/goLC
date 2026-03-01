package dp

import "testing"

func TestDecodings(t *testing.T) {

	c1 := "1201"
	if numDecodings(c1) != 1 {
		t.FailNow()
	}
	c2 := "2101"
	if numDecodings(c2) != 1 {
		t.FailNow()
	}

}
