package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetCommissionRateService struct {
	c      *Client
	symbol string
}

func (s *GetCommissionRateService) Symbol(symbol string) *GetCommissionRateService {
	s.symbol = symbol
	return s
}

func (s *GetCommissionRateService) Do(ctx context.Context, opts ...RequestOption) (res *CommissionRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/commissionRate",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(CommissionRate)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
