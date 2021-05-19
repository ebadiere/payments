package uniswap

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTwoHundredFiftySixDecimals(t *testing.T) {
	address := common.HexToAddress("0x0000000000000000000000000000000000000002")
	_, err := NewToken(Ropsten, address, 256, "dai", "dai")
	if err == nil {
		t.Logf("Token was created with 256 decimals! No error:, %s", err)
		t.Fail()
	}
}

func TestTokenEqualsDifferentAddresses(t *testing.T) {

	daiAddress := common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d")
	testStandardToken := common.HexToAddress("0x722dd3F80BAC40c951b51BdD28Dd19d435762180")

	dai, err := NewToken(Ropsten, daiAddress, 18, "dai", "dai")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	unkn, err := NewToken(Ropsten, testStandardToken, 18, "unkn", "unkn")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	if dai.Equals(*unkn) {
		t.Logf("Failed equality test, %s", err)
		t.Fail()
	}

}

func TestTokenEqualsDifferentNameSymbol(t *testing.T) {

	daiAddress := common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d")
	dai, err := NewToken(Ropsten, daiAddress, 18, "dai", "dai")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	unkn, err := NewToken(Ropsten, daiAddress, 18, "unkn", "unkn")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	if !dai.Equals(*unkn) {
		t.Logf("Failed equality test, %s", err)
		t.Fail()
	}

}

func TestTokenEqualsDifferentChainId(t *testing.T) {

	daiAddress := common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d")
	dai, err := NewToken(Ropsten, daiAddress, 18, "dai", "dai")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	unkn, err := NewToken(Kovan, daiAddress, 18, "unkn", "unkn")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	if dai.Equals(*unkn) {
		t.Logf("Failed equality test, %s", err)
		t.Fail()
	}

}
