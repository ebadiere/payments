package uniswap

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestTwoHundredFiftySixDecimalsFail(t *testing.T) {
	address := common.HexToAddress("0x0000000000000000000000000000000000000002")
	_, err := NewToken(Ropsten, address, 256, "dai", "dai")
	if err == nil {
		t.Logf("Token was created with 256 decimals! No error:, %s", err)
		t.Fail()
	}
}

func TestTokenEqualsDifferentAddressesFalse(t *testing.T) {

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

func TestTokenEqualsDifferentNameSymbolTrue(t *testing.T) {

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

func TestTokenEqualsDifferentChainIdFalse(t *testing.T) {

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

func TestTokenEqualsDifferentDecimalsTrue(t *testing.T) {

	daiAddress := common.HexToAddress("0xad6d458402f60fd3bd25163575031acdce07538d")
	dai, err := NewToken(Ropsten, daiAddress, 18, "dai", "dai")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	unkn, err := NewToken(Ropsten, daiAddress, 10, "unkn", "unkn")

	if err != nil {
		t.Logf("Failed to create Token, %s", err)
		t.Fail()
	}

	if !dai.Equals(*unkn) {
		t.Logf("Failed equality test, %s", err)
		t.Fail()
	}

}

func TestTokenEqualsSameAddressTrue(t *testing.T) {

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
