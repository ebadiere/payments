package uniswap

import "math/big"


type CurrenyAmount struct {

	currency 		Token
	decimalScale	big.Int
	numerator 		big.Int
	denominator 	big.Int
	quotient 		big.Int

}

var currency 		Token
var decimalScale	big.Int

var denominator 	big.Int
var numerator 		big.Int
var quotient 		big.Int

type fraction struct{
	numerator 		big.Int
	denominator 	big.Int
}


func NewCurrencyAmount(_currency Token, fract fraction) CurrenyAmount {

	denominator	:= fract.denominator

	if denominator.Cmp(big.NewInt(0)) <= 0 {
		denominator = *big.NewInt(1)
	}

	quotient = getQuotient(&numerator, &denominator)
	currency = _currency
	ten := big.NewInt(10)
	_decimals := big.NewInt( int64(_currency.baseCurrency.decimals))
	decimalScale = *ten.Exp(ten, _decimals, nil)

	currencyAmount := CurrenyAmount{
		currency: currency,
		decimalScale: decimalScale,
		numerator: fract.numerator,
		denominator: denominator,
		quotient: quotient,
	}

	return currencyAmount
}

func getQuotient(_numerator *big.Int, _denominator *big.Int) big.Int{

	q := new(big.Int)
	q.Div(_numerator, _denominator)
	return *q
}


