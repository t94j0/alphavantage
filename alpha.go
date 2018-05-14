package alphavantage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/tidwall/gjson"
)

const endpointURL = "https://www.alphavantage.co/query?"

// Config defines the configuration for an alphavantage data getter
type Config struct {
	apiKey string
}

// New creates a new alphavantage Config struct
func New(apiKey string) Config {
	return Config{apiKey}
}

// StockIntraday gets intraday stock information for the requested symbol
func (c Config) StockIntraday(
	interval Interval,
	symbol string,
) (o IntraDayAPI, err error) {
	values := url.Values{
		"function": {Intraday.Function},
		"symbol":   {symbol},
		"interval": {string(interval)},
	}

	body, err := c.getter(values)
	if err != nil {
		return o, err
	}

	var output IntraDayAPI

	metaData := gjson.GetBytes(body, "Meta Data")
	if err := json.Unmarshal([]byte(metaData.String()), &output.MetaData); err != nil {
		return o, err
	}

	stockData := gjson.GetBytes(body, "Time Series*")
	if err := json.Unmarshal([]byte(stockData.String()), &output.data); err != nil {
		return o, err
	}

	for t, val := range output.data {
		newTime, err := time.Parse(Intraday.TimeString, t)
		if err != nil {
			return o, err
		}
		val.Time = newTime
		output.Data = append(output.Data, val)
	}

	return output, nil
}

// Newest gets the newest event from the returned list of data points
func (id IntraDayAPI) Newest() (o DatapointAPI) {
	if len(id.Data) == 0 {
		return
	}

	index := 0
	var savedTime time.Time

	for i, data := range id.Data {
		if data.Time.After(savedTime) {
			savedTime = data.Time
			index = i
		}
	}

	return id.Data[index]
}

// getter gets data from the API
func (c Config) getter(values url.Values) (out []byte, err error) {
	values.Add("apikey", c.apiKey)
	query := values.Encode()
	fullURL := endpointURL + query

	response, err := http.Get(fullURL)
	if err != nil {
		return out, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return out, err
	}

	return data, nil
}
