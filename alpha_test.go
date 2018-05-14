package alphavantage

import (
	"fmt"
	"testing"
)

func ExampleConfig_StockIntraday() {
	// This test is replicating the example found in:
	// https://www.alphavantage.co/documentation/#intraday
	const DemoKey = "demo"
	config := New(DemoKey)

	out, _ := config.StockIntraday(OneMin, "MSFT")
	fmt.Println(out.MetaData.Interval)
	// output:
	// 1min
}

func ExampleIntraDayAPI_Newest() {
	const DemoKey = "demo"
	config := New(DemoKey)

	out, _ := config.StockIntraday(OneMin, "MSFT")
	dp := out.Newest()
	fmt.Println(dp.Close)
	// output:
	// 97.7000
}

func TestConfig_StockIntraday(t *testing.T) {
	config := New("demo")

	_, err := config.StockIntraday(OneMin, "MSFT")
	if err != nil {
		t.Fatal(err)
	}
}
