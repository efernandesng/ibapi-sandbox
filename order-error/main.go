package main

import (
	"fmt"

	"github.com/scmhub/ibsync"
)

func main() {
	log := ibsync.Logger()
	ibsync.SetLogLevel(0)
	ibsync.SetConsoleWriter()

	// New IB client & Connect
	ib := ibsync.NewIB()

	err := ib.Connect(
		ibsync.NewConfig(
			ibsync.WithHost("127.0.0.1"),
			ibsync.WithPort(7497),
			ibsync.WithClientID(10),
		),
	)
	if err != nil {
		log.Error().Err(err).Msg("Connect")
		return
	}
	defer ib.Disconnect()

	nvda := ibsync.NewStock("NVDA", "NASDAQ", "USD")
	order := ibsync.MarketOrder("BUY", ibsync.StringToDecimal("1"))
	trade := ib.PlaceOrder(nvda, order)

	<-trade.Done() // Will become stuck here if the order submission fails with an error
	// time.Sleep(5 * time.Second) // Wait for the trade to complete

	fmt.Printf("%d Trade Logs:\n", len(trade.Logs())) // No error logs present even if the order is rejected
	for _, log := range trade.Logs() {
		fmt.Printf("  - %#v\n", log)
	}
}
