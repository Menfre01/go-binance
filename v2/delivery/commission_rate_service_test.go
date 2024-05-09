package delivery

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type baseCommissionRateTestSuite struct {
	baseTestSuite
}

type commissionRateServiceTestSuite struct {
	baseCommissionRateTestSuite
}

func TestCommissionRateService(t *testing.T) {
	suite.Run(t, new(commissionRateServiceTestSuite))
}

func (s *commissionRateServiceTestSuite) TestGetCommissionRate() {
	data := []byte(`{
    "symbol": "BTCUSD_PERP",
    "makerCommissionRate": "0.00015",
    "takerCommissionRate": "0.00040"
	}`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_PERP"
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol": symbol,
		})
		s.assertRequestEqual(e, r)
	})
	order, err := s.client.NewGetCommissionRateService().Symbol(symbol).Do(newContext())
	r := s.r()
	r.NoError(err)
	e := &CommissionRate{
		Symbol:              "BTCUSD_PERP",
		MakerCommissionRate: "0.00015",
		TakerCommissionRate: "0.00040",
	}
	s.assertCommissionRateEqual(e, order)
}

func (s *commissionRateServiceTestSuite) assertCommissionRateEqual(e, a *CommissionRate) {
	r := s.r()
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.TakerCommissionRate, a.TakerCommissionRate, "TakerCommissionRate")
	r.Equal(e.MakerCommissionRate, a.MakerCommissionRate, "MakerCommissionRate")
}
