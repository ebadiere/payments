package uniswap

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
)

const Mainnet = 1
const Ropsten = 3
const Rinkeby = 4
const Gorli = 5
const Kovan = 42

type BaseCurrency struct {
	isEther bool
	isToken bool

	decimals int
	symbol   string
	name     string
}

type Token struct {
	baseCurrency BaseCurrency
	chainID      int
	address      common.Address
}

func NewToken(chainId int, address string, decimals int, symbol string, name string) (*Token, error) {
	base := BaseCurrency{
		isEther:  false,
		isToken:  true,
		decimals: decimals,
		symbol:   symbol,
		name:     name,
	}

	_, err := hex.DecodeString(address)

	return &Token{
		baseCurrency: base,
		chainID:      chainId,
		address:      common.HexToAddress(address),
	}, err
}
