package sentiment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type SentimentData struct {
	Value int
	Name  string
	Time  time.Time
}

type Sentiment []SentimentData

type sentimentBlob struct {
	Name string `json:"name"`
	Data []struct {
		Value     string `json:"value"`
		Name      string `json:"value_classification"`
		Time      string `json:"timestamp"`
		TimeUntil string `json:"time_until_update"`
	}
	Metadata interface{} `json:"metadata"`
}

// GetSentiment returns list of SentimentData for N(=limit) days
func GetSentiment(limit int) (Sentiment, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.alternative.me/fng/?limit=%d", limit))
	if err != nil || resp.StatusCode != 200 {
		return Sentiment{}, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Sentiment{}, err
	}
	resp.Body.Close()

	var blob sentimentBlob
	json.Unmarshal(data, &blob)

	var list = make(Sentiment, len(blob.Data))
	for i, v := range blob.Data {
		val, _ := strconv.Atoi(v.Value)
		ts, _ := strconv.Atoi(v.Time)
		list[i].Value = val
		list[i].Name = v.Name
		list[i].Time = time.Unix(int64(ts), 0)
	}

	return list, nil
}
