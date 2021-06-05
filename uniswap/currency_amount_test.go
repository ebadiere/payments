package uniswap

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var addressOne = common.HexToAddress("0x0000000000000000000000000000000000000001")

func TestCurrencyAmountTest(t *testing.T){

	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating Test token:, %s", err)
		t.Fail()
	}
	
	f := Fraction{numerator: *big.NewInt(100)}
	amount, err := NewCurrencyAmount(*token, f)
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}	

	t.Log("Quotient: %", amount.quotient)
	if amount.quotient.Cmp(big.NewInt(100)) != 0{
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
	amount, err := NewCurrencyAmount(*token, f)
	if err != nil {
		t.Logf("Error creating Test token:, %s", err)
		t.Fail()
	}	

	newCurencyAmount := amount.Multiply(percent)
	expectedResult := *big.NewInt(15)

	if newCurencyAmount.quotient.Cmp(&expectedResult) != 0 {
		t.Logf("Error comparing quotient results:, %s", err)		
		t.Logf("Expected expectedResult:, %s", &expectedResult)		
		t.Logf("Actual result:, %s", &newCurencyAmount.quotient)
		t.Fail()		
	}

}
func TestProducesEtherAmount(t *testing.T){
	ether, err := NewEther(Mainnet)
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}

	f := Fraction{numerator: *big.NewInt(100)}

	amount, err := NewCurrencyAmount(*ether, f)
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}	

	expectedResult := *big.NewInt(100)

	if amount.quotient.Cmp(&expectedResult) != 0 {
		t.Logf("Error comparing quotient results:, %s", err)		
		t.Logf("Expected expectedResult:, %s", &expectedResult)		
		t.Logf("Actual result:, %s", &amount.quotient)
		t.Fail()		
	}

	if amount.currency.chainID != ether.chainID {
		t.Logf("Error comparing chainId results:, %s", err)		
		t.Logf("Expected expectedResult:, %d", ether.chainID)		
		t.Logf("Actual result:, %d", &amount.currency.chainID)
		t.Fail()		
	}	
}

func TestAmountCanBeMaxUint256(t *testing.T){
	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}
	
	f := Fraction{numerator: *big.NewInt(maxUint256.Hash().Big().Int64())}
	amount, err := NewCurrencyAmount(*token, f)
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}	

	t.Log("Quotient: %", amount.quotient)
	if amount.quotient.Cmp(big.NewInt(maxUint256.Hash().Big().Int64())) != 0{
		t.Errorf("Incorrect quotient")
	}
}

func TestAmountCannotExceedUint256(t *testing.T){
	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}
	
	numerator, err := hexutil.DecodeBig("0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff")
	if err != nil {
		t.Logf("Error creating max uint256 value:, %s", err)
		t.Fail()
	}

	// numerator := big.NewInt(maxUint256.Hash().Big().Int64())

	numerator = numerator.Add(numerator, big.NewInt(1))

	f := Fraction{numerator: *numerator}
	amount, err := NewCurrencyAmount(*token, f)
	if err == nil {
		t.Logf("Error CurrencyAmount should not have been created:, %s", err)
		t.Fail()
	}	
	t.Log("Quotient: %", amount.quotient)
}

func TestNumeratorCanBeGTIfDenominatorGTOne(t *testing.T){

	token, err := NewToken(Mainnet, addressOne, 18, "tst", "Test")
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}

	numerator := big.NewInt(maxUint256.Hash().Big().Int64())

	numerator = numerator.Add(numerator, big.NewInt(2))
	denominator := big.NewInt(2)

	f := Fraction{numerator: *numerator, denominator: *denominator}

	amount, err := NewCurrencyAmount(*token, f)
	if err != nil {
		t.Logf("Error creating CurrencyAmount:, %s", err)
		t.Fail()
	}

	t.Log("Quotient: %", amount.quotient)

}

