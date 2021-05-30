package uniswap

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var addressOne = common.HexToAddress("0x0000000000000000000000000000000000000001")

func TestCurrencyAmountTest(t *testing.T){

	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating Test token:, %s", err)
		t.Fail()
	}
	
	f := Fraction{numerator: *big.NewInt(100)}
	amount := NewCurrencyAmount(*token, f)
	if amount.quotient.Cmp(big.NewInt(100)) == 0{
		t.Errorf("Incorrect quotient")
	}

}

func TestReturnsAmountAfterMultiplication(t *testing.T){
	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating Test token:, %s", err)
		t.Fail()
	}	

	f := Fraction{numerator: *big.NewInt(100)}
	percent := Percent{
		numerator: 		*big.NewInt(15),
		denominator: 	*big.NewInt(100),
	}
	amount := NewCurrencyAmount(*token, f)

	newCurencyAmount := amount.Multiply(percent)
	expectedResult := *big.NewInt(15)

	if newCurencyAmount.quotient.Cmp(&expectedResult) != 0 {
		t.Logf("Error comparing quotient results:, %s", err)		
		t.Logf("Expected expectedResult:, %s", &expectedResult)		
		t.Logf("Actual result:, %s", &newCurencyAmount.quotient)
		t.Fail()		
	}

}


