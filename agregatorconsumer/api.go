package agregatorconsumer

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	aggregatorv3 "github.com/nikola43/chainlinkpricefeed/aggregatorv3"
)

func (pfc *PriceFeedConsumer) GetPrice(feedAddress string) (*PriceFeedData, error) {

	// Test if it is a contract address.
	ok := IsContractAddress(feedAddress, pfc.client)
	if !ok {
		log.Fatalf("address %s is not a contract address\n", feedAddress)
	}

	chainlinkPriceFeedProxyAddress := common.HexToAddress(feedAddress)
	chainlinkPriceFeedProxy, err := aggregatorv3.NewAggregatorV3Interface(chainlinkPriceFeedProxyAddress, pfc.client)
	if err != nil {
		log.Fatal(err)
	}

	roundData, err := chainlinkPriceFeedProxy.LatestRoundData(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := chainlinkPriceFeedProxy.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	description, err := chainlinkPriceFeedProxy.Description(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	// Compute a big.int which is 10**decimals.
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)

	return &PriceFeedData{
		FeedAddress:     feedAddress,
		RoundId:         roundData.RoundId,
		Answer:          roundData.Answer,
		StartedAt:       roundData.StartedAt,
		UpdatedAt:       roundData.UpdatedAt,
		AnsweredInRound: roundData.AnsweredInRound,
		Decimals:        decimals,
		Divisor:         divisor,
		Description:     description,
	}, nil
}
