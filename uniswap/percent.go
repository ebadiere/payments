package uniswap

import "math/big"



type Percent struct{
	numerator 		big.Int
	denominator 	big.Int
}


func NewPercent(frac Fraction) Percent {
	percent := Percent {
		numerator: 		frac.numerator,
		denominator: 	frac.denominator,
	}

	return percent
}

func (percent Percent) ToDecimal() *big.Int {
	dv := new(big.Int).Div(&percent.numerator, &percent.denominator)
	return dv
}