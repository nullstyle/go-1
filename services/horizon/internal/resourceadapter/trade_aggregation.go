package resourceadapter

import (
	"context"

	"github.com/stellar/go/amount"
	"github.com/stellar/go/price"
	protocol "github.com/stellar/go/protocols/horizon"
	"github.com/stellar/go/services/horizon/internal/db2/history"
)

// Populate fills out the details of a trade using a row from the history_trades
// table.
func PopulateTradeAggregation(
	ctx context.Context,
	dest *protocol.TradeAggregation,
	row history.TradeAggregation,
) error {
	var err error
	dest.Timestamp = row.Timestamp
	dest.TradeCount = row.TradeCount
	dest.BaseVolume, err = amount.IntStringToAmount(row.BaseVolume)
	if err != nil {
		return err
	}
	dest.CounterVolume, err = amount.IntStringToAmount(row.CounterVolume)
	if err != nil {
		return err
	}
	dest.Average = price.StringFromFloat64(row.Average)
	dest.High = row.High.XDR.String()
	dest.HighR = row.High.XDR
	dest.Low = row.Low.XDR.String()
	dest.LowR = row.Low.XDR
	dest.Open = row.Open.XDR.String()
	dest.OpenR = row.Open.XDR
	dest.Close = row.Close.XDR.String()
	dest.CloseR = row.Close.XDR
	return nil
}
