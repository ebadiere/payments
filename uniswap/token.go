package uniswap

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
)


type BaseCurrency struct {
	isEther bool
	isToken bool

	decimals int
	symbol   string
	name     string
}

type Token struct {
	baseCurrency BaseCurrency
	chainID      uint
	address      common.Address
}

func NewToken(chainId uint, address common.Address, decimals int, symbol string, name string) (*Token, error) {
	base := BaseCurrency{
		isEther:  false,
		isToken:  true,
		decimals: decimals,
		symbol:   symbol,
		name:     name,
	}

	if decimals > 255 {
		return &Token{}, errors.New("DECIMALS")
	}

	return &Token{
		baseCurrency: base,
		chainID:      chainId,
		address:      address,
	}, nil
}

/**
* Returns true if the tokens have the same chainId and address
 */
func (t Token) Equals(u Token) bool {

	if (t.address == u.address) && (t.chainID == u.chainID) {
		return true
	}

	return false

}
