package uniswap

type Ether struct {
	baseCurrency BaseCurrency
	chainID      uint
}

func NewEther(chainId uint) (*Ether, error) {
	base := BaseCurrency{
		isEther:  true,
		isToken:  false,
		decimals: 18,
		symbol:   "ETH",
		name:     "Ether",
	}

	return &Ether{
		baseCurrency: base,
		chainID:      chainId,
	}, nil
}