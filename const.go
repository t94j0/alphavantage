package alphavantage

import "time"

// Interval defines intervals for Intraday
type Interval string

const (
	// OneMin gets 1-minute difference records
	OneMin Interval = "1min"

	// FiveMin gets 5-minute difference records
	FiveMin Interval = "5min"

	// FifteenMin gets 15-minute difference records
	FifteenMin Interval = "15min"

	// ThirtyMin gets 30-minute difference records
	ThirtyMin Interval = "30min"

	// SixtyMin gets 60-minute difference records
	SixtyMin Interval = "60min"
)

// TimeSeries is used for determining which time series should be used for the
// curent API request
type TimeSeries struct {
	Function   string
	TimeString string
}

var (
	// Intraday defines the TimeSeries for intra-day requests
	Intraday = TimeSeries{"TIME_SERIES_INTRADAY", "2006-01-02 15:04:05"}
	/*Daily           = TimeSeries{"TIME_SERIES_DAILY", "Daily"}
	DailyAdjusted   = TimeSeries{"TIME_SERIES_DAILY_ADJUSTED", "Daily"}
	Weekly          = TimeSeries{"TIME_SERIES_WEEKLY", "Weekly"}
	WeeklyAdjusted  = TimeSeries{"TIME_SERIES_WEEKLY_ADJUSTED", "Weekly"}
	Monthly         = TimeSeries{"TIME_SERIES_MONTHLY", "Monthly"}
	MonthlyAdjusted = TimeSeries{"TIME_SERIES_MONTHLY_ADJUSTED", "Monthly"}
	*/
)

// IntraDayAPI is a struct used for returning API data to
type IntraDayAPI struct {
	MetaData MetaData
	Data     []DatapointAPI

	data map[string]DatapointAPI
}

// MetaData is where meta data is marshalled to
type MetaData struct {
	Information string `json:"1. Information"`
	Symbol      string `json:"2. Symbol"`
	LastRefresh string `json:"3. Last Refreshed"`
	Interval    string `json:"4. Interval"`
	OutputSize  string `json:"5. Output Size"`
	TimeZone    string `json:"6. Time Zone"`
}

// DatapointAPI is where data point API data is marshalled into
// BUG(t94j0): Convert strings to actual data values (floats)
type DatapointAPI struct {
	Time   time.Time
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}
