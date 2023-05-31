package main

import (
	"fmt"

	"github.com/slicken/sentiment"
)

func main() {
	// get latest sentimentData
	// replace 0 with number of days you want data from
	sentimentData, err := sentiment.GetSentiment(0)
	if err != nil {
		panic(err)
	}

	fmt.Printf("SentimentDate  %v\nSentimentValue %v\nSentimentName  %v\n",
		sentimentData[0].Time,
		sentimentData[0].Value,
		sentimentData[0].Name)
}
