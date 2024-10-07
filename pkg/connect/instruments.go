package mbconnect

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Instrument struct {
	InstrumentToken uint32  `json:"instrument_token"`
	ExchangeToken   uint32  `json:"exchange_token"`
	Tradingsymbol   string  `json:"tradingsymbol"`
	Name            string  `json:"name"`
	LastPrice       float64 `json:"last_price"`
	Expiry          string  `json:"expiry"`
	Strike          float64 `json:"strike"`
	TickSize        float64 `json:"tick_size"`
	LotSize         uint    `json:"lot_size"`
	InstrumentType  string  `json:"instrument_type"`
	Segment         string  `json:"segment"`
	Exchange        string  `json:"exchange"`
}

type InstrumentsQueryParams struct {
	Exchange        string
	Tradingsymbol   string
	InstrumentToken uint32
	Name            string
	Expiry          string
	Strike          float64
	Segment         string
	InstrumentType  string
}

type OptionChainQueryParams struct {
	Exchange  string
	Name      string
	FutExpiry string
	OptExpiry string
}

// -----------------------------------------------------
// POST /instruments/info?s=NSE:NIFTY%2050&s=BSE:SENSEX
// -----------------------------------------------------
func (c *Client) InstrumentsInfoBySymbols(symbols []string) (map[string]Instrument, error) {
	fmt.Println("symbols len", len(symbols))
	if len(symbols) == 0 {
		return nil, fmt.Errorf("`symbols` are required")
	}
	params := url.Values{}
	for _, symbol := range symbols {
		params.Add("s", symbol)
	}
	var symbolInstrumentMap map[string]Instrument
	err := c.doEnvelope(http.MethodGet, URIInstrumentsInfo, params, nil, &symbolInstrumentMap)
	return symbolInstrumentMap, err
}

// -----------------------------------------------------
// GET /instruments/info?t=256265&t=8961794
// -----------------------------------------------------
func (c *Client) InstrumentsInfoByTokens(tokens []uint32) (map[uint32]Instrument, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("`tokens` are required")
	}
	stringTokens := make([]string, len(tokens))
	for i, token := range tokens {
		stringTokens[i] = strconv.FormatUint(uint64(token), 10)
	}
	params := url.Values{}
	for _, stringToken := range stringTokens {
		params.Add("t", stringToken)
	}
	var tokenInstruments map[uint32]Instrument
	err := c.doEnvelope(http.MethodGet, URIInstrumentsInfo, params, nil, &tokenInstruments)
	return tokenInstruments, err
}

// -----------------------------------------------------
// GET /instruments/query?exchange=NSE&tradingsymbol=SBIN
// -----------------------------------------------------
func (c *Client) InstrumentsQuery(qp InstrumentsQueryParams) ([]Instrument, error) {
	params := makeQueryParams(qp)
	var instruments []Instrument
	err := c.doEnvelope(http.MethodGet, URIInstrumentsQuery, params, nil, &instruments)
	return instruments, err
}

// makeQueryParams makes query params - helper function
func makeQueryParams(qp InstrumentsQueryParams) url.Values {
	params := url.Values{}
	if qp.Exchange != "" {
		params.Add("exchange", qp.Exchange)
	}
	if qp.Tradingsymbol != "" {
		params.Add("tradingsymbol", qp.Tradingsymbol)
	}
	if qp.InstrumentToken != 0 {
		params.Add("instrument_token", strconv.FormatUint(uint64(qp.InstrumentToken), 10))
	}
	if qp.Name != "" {
		params.Add("name", qp.Name)
	}
	if qp.Expiry != "" {
		params.Add("expiry", qp.Expiry)
	}
	if qp.Strike != 0 {
		params.Add("strike", strconv.FormatFloat(qp.Strike, 'f', -1, 64))
	}
	if qp.Segment != "" {
		params.Add("segment", qp.Segment)
	}
	if qp.InstrumentType != "" {
		params.Add("instrument_type", qp.InstrumentType)
	}
	return params
}

// -----------------------------------------------------
// GET /instruments/optionchain?exchange=NFO&name=NIFTY&opt_expiry=2024-10-10&fut_expiry=2024-10-31
// -----------------------------------------------------
func (c *Client) OptionchainInstruments(ocp OptionChainQueryParams) ([]Instrument, error) {
	params, err := makeOptionChainParams(ocp)
	if err != nil {
		return nil, err
	}
	var instruments []Instrument
	err = c.doEnvelope(http.MethodGet, URIInstrumentsOptionchain, params, nil, &instruments)
	return instruments, err
}

// OptionchainTokensMap gets option chain tokens to symbols map
func (c *Client) OptionchainTokenSymbolMap(ocp OptionChainQueryParams) (map[uint32]string, error) {
	instruments, err := c.OptionchainInstruments(ocp)
	if err != nil {
		return nil, err
	}
	tokenSymbolMap := make(map[uint32]string)
	for _, instrument := range instruments {
		tokenSymbolMap[instrument.InstrumentToken] = instrument.Exchange + ":" + instrument.Tradingsymbol
	}
	return tokenSymbolMap, nil
}

// makeOptionChainParams makes query params for option chain - helper function
func makeOptionChainParams(ocp OptionChainQueryParams) (url.Values, error) {
	params := url.Values{}
	if ocp.Exchange != "" {
		params.Add("exchange", ocp.Exchange)
	}
	if ocp.Name != "" {
		params.Add("name", ocp.Name)
	}
	if ocp.FutExpiry != "" {
		params.Add("fut_expiry", ocp.FutExpiry)
	}
	if ocp.OptExpiry != "" {
		params.Add("opt_expiry", ocp.OptExpiry)
	}
	return params, nil
}

// -----------------------------------------------------
// GET /instruments/fno/segment_expiries/:name
// -----------------------------------------------------
func (c *Client) FNOSegmentExpiries(name string) (map[string][]string, error) {
	if name == "" {
		return nil, fmt.Errorf("`name` is required")
	}
	var segmentExpiriesMap map[string][]string
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIInstrumentsFNOSegmentExpiries, name), nil, nil, &segmentExpiriesMap)
	return segmentExpiriesMap, err
}

// -----------------------------------------------------
// GET /instruments/fno/segment_expiries/:name
// -----------------------------------------------------
func (c *Client) FNOSegmentNames(expiry string) (map[string][]string, error) {
	if expiry == "" {
		return nil, fmt.Errorf("`expiry` is required")
	}
	var segmentNamesMap map[string][]string
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIInstrumentsFNOSegmentNames, expiry), nil, nil, &segmentNamesMap)
	return segmentNamesMap, err
}
