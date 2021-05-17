package uniswap

import (
	"testing"
)

func TestInvalidAddress(t *testing.T) {

	_, err := NewToken(Ropsten, "0xhello00000000000000000000000000000000002", 18, "dai", "dai")
	if err == nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

}
