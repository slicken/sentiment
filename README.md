# sentiment

crypto sentiment / fear and greed index<br>
very simple package, gets sentiment data from https://alternative.me

```
// get latest sentimentData
// replace 0 with number of days you want data for
sentimentData, err := sentiment.GetSentiment(1)
if err != nil {
    panic(err)
}

fmt.Printf("SentimentDate  %v\nSentimentValue %v\nSentimentName  %v\n",
    sentimentData[0].Time,
    sentimentData[0].Value,
    sentimentData[0].Name)
```
