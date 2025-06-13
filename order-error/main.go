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

	cccmf := ibsync.NewStock("CCCMF", "SMART", "USD")
	order := ibsync.MarketOrder("BUY", ibsync.StringToDecimal("1"))
	trade := ib.PlaceOrder(cccmf, order)

	<-trade.Done()

	fmt.Printf("%d Trade Logs:\n", len(trade.Logs()))
	for _, log := range trade.Logs() {
		fmt.Printf("  - %#v\n", log)
	}
}
