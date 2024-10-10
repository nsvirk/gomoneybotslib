package mbconnect

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Instrument is a struct that represents a trading instrument
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

// InstrumentsQueryParams is a struct that represents the query parameters for instruments
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

// POST /instruments/info?s=NSE:NIFTY%2050&s=BSE:SENSEX - Get instruments info by `symbols`
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

// GET /instruments/info?t=256265&t=8961794 - Get instruments info by tokens
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

// GET /instruments/query?exchange=NSE&tradingsymbol=SBIN - Get instruments by query params
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

// GET /instruments/fno/segment_expiries/:name - Get FNO segment expiries by `name`
func (c *Client) FNOSegmentExpiries(name string) (map[string][]string, error) {
	if name == "" {
		return nil, fmt.Errorf("`name` is required")
	}
	var segmentExpiriesMap map[string][]string
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIInstrumentsFNOSegmentExpiries, name), nil, nil, &segmentExpiriesMap)
	return segmentExpiriesMap, err
}

// GET /instruments/fno/segment_expiries/:name - Get FNO segment names by expiry
func (c *Client) FNOSegmentNames(expiry string) (map[string][]string, error) {
	if expiry == "" {
		return nil, fmt.Errorf("`expiry` is required")
	}
	var segmentNamesMap map[string][]string
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIInstrumentsFNOSegmentNames, expiry), nil, nil, &segmentNamesMap)
	return segmentNamesMap, err
}
