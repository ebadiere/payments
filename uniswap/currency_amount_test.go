package uniswap

import (
	"math/big"
	"testing"
	"github.com/ethereum/go-ethereum/common"
)


func TestCurrencyAmountTest(t *testing.T){

	addressOne := common.HexToAddress("0x0000000000000000000000000000000000000001")

	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating Test token:, %s", err)
		t.Fail()
	}
	
	f := fraction{numerator: *big.NewInt(100)}
	amount := NewCurrencyAmount(*token, f)
	if amount.quotient.Cmp(big.NewInt(100)) == 0{
		t.Errorf("Incorrect quotient")
	}

}


