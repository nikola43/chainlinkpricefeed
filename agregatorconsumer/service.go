package agregatorconsumer

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type PriceFeedData struct {
	FeedAddress     string
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
	Decimals        uint8
	Divisor         *big.Int
	Description     string
}

type PriceFeedConsumer struct {
	client *ethclient.Client
}

func New(rpcUrl string) *PriceFeedConsumer {
	c, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatal(err)
	}
	return &PriceFeedConsumer{client: c}
}
