package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListUserTradesService struct {
	c         *Client
	symbol    string
	pair      string
	orderID   *int64
	startTime *int64
	endTime   *int64
	limit     *int
	fromId    *int64
}

func (s *ListUserTradesService) Symbol(symbol string) *ListUserTradesService {
	s.symbol = symbol
	return s
}

// Pair set pair
func (s *ListUserTradesService) Pair(pair string) *ListUserTradesService {
	s.pair = pair
	return s
}

// OrderID set orderID
func (s *ListUserTradesService) OrderID(orderID int64) *ListUserTradesService {
	s.orderID = &orderID
	return s
}

// StartTime set starttime
func (s *ListUserTradesService) StartTime(startTime int64) *ListUserTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListUserTradesService) EndTime(endTime int64) *ListUserTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListUserTradesService) Limit(limit int) *ListUserTradesService {
	s.limit = &limit
	return s
}

func (s *ListUserTradesService) FromId(fromId int64) *ListUserTradesService {
	s.fromId = &fromId
	return s
}

// Do send request
func (s *ListUserTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*Trade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/userTrades",
		secType:  secTypeSigned,
	}
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	if s.pair != "" {
		r.setParam("pair", s.pair)
	}
	if s.orderID != nil {
		r.setParam("orderId", *s.orderID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.fromId != nil {
		r.setParam("fromId", *s.fromId)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Trade{}, err
	}
	res = make([]*Trade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*Trade{}, err
	}
	return res, nil
}

type Trade struct {
	Symbol          string `json:"symbol"`
	ID              int64  `json:"id"`
	OrderID         int64  `json:"orderId"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	MarginAsset     string `json:"marginAsset"`
	BaseQty         string `json:"baseQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	PositionSide    string `json:"positionSide"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
}
