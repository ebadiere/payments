package uniswap

import (
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

var currencyMap map[string]interface{}

func NewCurrency(chainId uint, param ...string) (*Token, error) {

	if len(param) == 1 {
		return NewEther(chainId)
	}

	if len(param) != 5 {
		log.Fatal("Incorrect number of currency parameters")
	}

	address := common.HexToAddress(param[1])

	decimals, err := strconv.ParseUint(param[2], 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	symbol := param[3]
	name := param[4]

	return NewToken(chainId, address, uint(decimals), symbol, name)

}

func NewEther(chainId uint) (*Token, error) {
	base := BaseCurrency{
		isEther:  true,
		isToken:  false,
		decimals: 18,
		symbol:   "ETH",
		name:     "Ether",
	}

	return &Token{
		baseCurrency: base,
		chainID:      chainId,
		address:      common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"),
	}, nil
}
