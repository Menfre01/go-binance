package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseTradeTestSuite struct {
	baseTestSuite
}

type tradeServiceTestSuite struct {
	baseTradeTestSuite
}

func TestTradeService(t *testing.T) {
	suite.Run(t, new(tradeServiceTestSuite))
}

func (s *tradeServiceTestSuite) TestListUserTrades() {
	data := []byte(`[
		{
			"symbol": "BTCUSD_200626",     
			"id": 6,                       
			"orderId": 28,                 
			"pair": "BTCUSD",              
			"side": "SELL",                
			"price": "8800",              
			"qty": "1",                   
			"realizedPnl": "0",            
			"marginAsset": "BTC",          
			"baseQty": "0.01136364",       
			"commission": "0.00000454", 
			"commissionAsset": "BTC",     
			"time": 1590743483586,        
			"positionSide": "BOTH",       
			"buyer": false,                
			"maker": false                  
		}
	  ]`)
	s.mockDo(data, nil)
	defer s.assertDo()
	symbol := "BTCUSD_200626"
	orderID := int64(28)
	limit := 3
	startTime := int64(1590743483586)
	endTime := int64(1590743483587)
	fromId := int64(0)
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"orderId":   orderID,
			"startTime": startTime,
			"endTime":   endTime,
			"limit":     limit,
			"fromId":    0,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewListUserTradesService().Symbol(symbol).
		OrderID(orderID).StartTime(startTime).EndTime(endTime).
		FromId(fromId).
		Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)

	e := &Trade{
		Symbol:          "BTCUSD_200626",
		ID:              6,
		OrderID:         28,
		Pair:            "BTCUSD",
		Side:            "SELL",
		Price:           "8800",
		Qty:             "1",
		RealizedPnl:     "0",
		MarginAsset:     "BTC",
		BaseQty:         "0.01136364",
		Commission:      "0.00000454",
		CommissionAsset: "BTC",
		Time:            1590743483586,
		PositionSide:    "BOTH",
		Buyer:           false,
		Maker:           false,
	}
	s.assertTradeEqual(e, trades[0])
}

func (s *baseTradeTestSuite) assertTradeEqual(e, a *Trade) {
	r := s.r()
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Qty, a.Qty, "Qty")
	r.Equal(e.RealizedPnl, a.RealizedPnl, "RealizedPnl")
	r.Equal(e.MarginAsset, a.MarginAsset, "MarginAsset")
	r.Equal(e.BaseQty, a.BaseQty, "BaseQty")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Buyer, a.Buyer, "Buyer")
	r.Equal(e.Maker, a.Maker, "Maker")
}
