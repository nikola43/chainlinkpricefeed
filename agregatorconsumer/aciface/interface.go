package pfiface

import (
	agregatorconsumer "github.com/nikola43/chainlinkpricefeed/agregatorconsumer"
)

type PriceFeedConsumerApi interface {
	GetPrice(feedAddress string) (agregatorconsumer.PriceFeedData, error)
}
