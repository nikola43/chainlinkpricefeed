package main

import (
	"github.com/nikola43/chainlinkpricefeed/agregatorconsumer"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// Read the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the rpc_url.
	rpcUrl := os.Getenv("RPC_URL")
	if len(rpcUrl) == 0 {
		log.Fatal("rpcUrl is empty. check the .env file")
	}

	// Assign default values to feedAddress, and update value if a feed address was passed in the command line.
	feedAddress := os.Getenv("DEFAULT_FEED_ADDR")
	if len(os.Args) > 1 && os.Args[1] != "" {
		feedAddress = os.Args[1]
	}

	pfc := agregatorconsumer.New(rpcUrl)
	pfd, err := pfc.GetPrice(feedAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v Price feed address is  %v\n", pfd.Description, feedAddress)
	log.Printf("Round id is %v\n", pfd.RoundId)
	log.Printf("Answer is %v\n", pfd.Answer)
	log.Printf("Formatted answer is %v\n", agregatorconsumer.DivideBigInt(pfd.Answer, pfd.Divisor))
	log.Printf("Started at %v\n", agregatorconsumer.FormatTime(pfd.StartedAt))
	log.Printf("Updated at %v\n", agregatorconsumer.FormatTime(pfd.UpdatedAt))
	log.Printf("Answered in round %v\n", pfd.AnsweredInRound)

}
