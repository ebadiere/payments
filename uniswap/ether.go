package uniswap

import "math/big"

const Mainnet = 1
const Ropsten = 3
const Rinkeby = 4
const Gorli = 5
const Kovan = 42

type Ether struct {

	currency 		Token
	decimalScale	big.Int
	numerator 		big.Int
	denominator 	big.Int
	quotient 		big.Int

}

func NewEther(chainId uint) (ether *Token, err error) {

	return NewNativeCurrency(chainId, 18, "Eth", "Ether")

}