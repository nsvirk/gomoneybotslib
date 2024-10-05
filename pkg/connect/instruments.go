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
	key             string
}

type OptionChainQueryParams struct {
	Exchange  string
	Name      string
	FutExpiry string
	OptExpiry string
	key       string
}

// InstrumentsBySymbols gets instruments by symbols
func (c *Client) InstrumentsBySymbols(symbols []string) (map[string]Instrument, error) {
	params := url.Values{}
	for _, symbol := range symbols {
		params.Add("s", symbol)
	}

	var symbolInstrumentMap map[string]Instrument
	err := c.doEnvelope(http.MethodGet, URIInstrumentsInfo, params, nil, &symbolInstrumentMap)
	return symbolInstrumentMap, err
}

// InstrumentsByTokens gets instruments by tokens
func (c *Client) InstrumentsByTokens(tokens []uint32) (map[uint32]Instrument, error) {
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

// TokensBySymbols gets tokens by symbols
func (c *Client) TokensBySymbols(symbols []string) ([]uint32, error) {
	symbolInstrumentMap, err := c.InstrumentsBySymbols(symbols)
	if err != nil {
		return nil, err
	}
	var tokens = make([]uint32, 0, len(symbolInstrumentMap))
	for _, symbolInstrumen := range symbolInstrumentMap {
		tokens = append(tokens, (symbolInstrumen.InstrumentToken))
	}
	return tokens, nil
}

// SymbolsByTokens gets symbols by tokens
func (c *Client) SymbolsByTokens(tokens []uint32) ([]string, error) {
	tokenInstrumentMap, err := c.InstrumentsByTokens(tokens)
	if err != nil {
		return nil, err
	}
	var symbols = make([]string, 0, len(tokenInstrumentMap))
	for _, instrument := range tokenInstrumentMap {
		symbols = append(symbols, (instrument.Exchange + ":" + instrument.Tradingsymbol))
	}
	return symbols, nil
}

// InstrumentsQuery gets instruments by query
func (c *Client) InstrumentsByQuery(qp InstrumentsQueryParams) (map[string]Instrument, error) {
	qp.key = ""
	params := makeQueryParams(qp)
	var instrumentsMap map[string]Instrument
	err := c.doEnvelope(http.MethodGet, URIInstrumentsQuery, params, nil, &instrumentsMap)
	return instrumentsMap, err
}

// SymbolsByQuery gets instrument symbols by query
func (c *Client) SymbolsByQuery(qp InstrumentsQueryParams) ([]string, error) {
	qp.key = "symbols"
	params := makeQueryParams(qp)
	var symbols []string
	err := c.doEnvelope(http.MethodGet, URIInstrumentsQuery, params, nil, &symbols)
	return symbols, err
}

// TokensByQuery gets instrument tokens by query
func (c *Client) TokensByQuery(qp InstrumentsQueryParams) ([]uint32, error) {
	qp.key = "tokens"
	params := makeQueryParams(qp)
	var stringTokens []string
	err := c.doEnvelope(http.MethodGet, URIInstrumentsQuery, params, nil, &stringTokens)
	var tokens []uint32
	for _, stringToken := range stringTokens {
		token, err := strconv.ParseUint(stringToken, 10, 32)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, uint32(token))
	}
	return tokens, err
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
	if qp.key != "" {
		params.Add("key", qp.key)
	}
	return params
}

// OptionchainInstruments gets option chain instruments
func (c *Client) OptionchainInstruments(ocp OptionChainQueryParams) (map[string]Instrument, error) {
	ocp.key = ""
	params, err := makeOptionChainParams(ocp)
	if err != nil {
		return nil, err
	}
	var optionChainMap map[string]Instrument
	err = c.doEnvelope(http.MethodGet, URIInstrumentsOptionchain, params, nil, &optionChainMap)
	return optionChainMap, err
}

// OptionchainSymbols gets option chain symbols
func (c *Client) OptionchainSymbols(ocp OptionChainQueryParams) ([]string, error) {
	ocp.key = "symbols"
	params, err := makeOptionChainParams(ocp)
	if err != nil {
		return nil, err
	}
	var symbols []string
	err = c.doEnvelope(http.MethodGet, URIInstrumentsOptionchain, params, nil, &symbols)
	return symbols, err
}

// OptionchainTokens gets option chain tokens
func (c *Client) OptionchainTokens(ocp OptionChainQueryParams) ([]uint32, error) {
	ocp.key = "tokens"
	params, err := makeOptionChainParams(ocp)
	if err != nil {
		return nil, err
	}
	var stringTokens []string
	err = c.doEnvelope(http.MethodGet, URIInstrumentsOptionchain, params, nil, &stringTokens)
	var tokens []uint32
	for _, stringToken := range stringTokens {
		token, err := strconv.ParseUint(stringToken, 10, 32)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, uint32(token))
	}
	return tokens, err
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
	if ocp.key != "" {
		params.Add("key", ocp.key)
	}
	return params, nil
}

// FNOSegmentExpiries gets instruments by FNO segment and name on expiry
func (c *Client) FNOSegmentExpiries(name string) (map[string][]string, error) {
	var segmentExpiriesMap map[string][]string
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIInstrumentsFNOSegmentExpiries, name), nil, nil, &segmentExpiriesMap)
	return segmentExpiriesMap, err
}

// FNOSegmentNames gets instruments by FNO segment and name on expiry
func (c *Client) FNOSegmentNames(expiry string) (map[string][]string, error) {
	var segmentNamesMap map[string][]string
	err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIInstrumentsFNOSegmentNames, expiry), nil, nil, &segmentNamesMap)
	return segmentNamesMap, err
}
