package mbconnect

import (
	"fmt"
	"net/http"
	"time"
)

type Index struct {
	ID            uint32    `json:"-"`
	Index         string    `json:"index"`
	Exchange      string    `json:"exchange"`
	Tradingsymbol string    `json:"tradingsymbol"`
	CompanyName   string    `json:"company_name"`
	Industry      string    `json:"industry"`
	Series        string    `json:"series"`
	ISINCode      string    `json:"isin_code"`
	UpdatedAt     time.Time `json:"-"`
}

// IndicesAll returns all the indices
func (c *Client) IndicesAll() (map[string][]Index, error) {
	var index = make(map[string][]Index)
	if err := c.doEnvelope(http.MethodGet, URIIndicesAll, nil, nil, &index); err != nil {
		return nil, fmt.Errorf("failed to get index names: %w", err)
	}
	return index, nil
}

// IndicesByExchange returns the indices for a given exchange
func (c *Client) IndicesByExchange(exchange string) ([]Index, error) {
	var index []Index
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesByExchange, exchange), nil, nil, &index); err != nil {
		return nil, fmt.Errorf("failed to get index names: %w", err)
	}
	return index, nil
}

// IndexNames returns the names of the indices for a given exchange
func (c *Client) IndexNames(exchange, name string) ([]string, error) {
	var indexNames []string
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesIndexNames, exchange, name), nil, nil, &indexNames); err != nil {
		return nil, fmt.Errorf("failed to get index names: %w", err)
	}
	return indexNames, nil
}

// IndexTokens returns the tokens of the indices for a given index name
func (c *Client) IndexTokens(exchange, name string) ([]string, error) {
	var tokens []string
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesIndexTokens, exchange, name), nil, nil, &tokens); err != nil {
		return nil, fmt.Errorf("failed to get index tokens: %w", err)
	}
	return tokens, nil
}

// IndexSymbols returns the symbols of the indices for a index name
func (c *Client) IndexSymbols(exchange, name string) ([]string, error) {
	var symbols []string
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesIndexSymbols, exchange, name), nil, nil, &symbols); err != nil {
		return nil, fmt.Errorf("failed to get index symbols: %w", err)
	}
	return symbols, nil
}

// IndexInstruments returns the instruments of the indices for a Index name
func (c *Client) IndexInstruments(exchange, name string) (map[string]Index, error) {
	var instruments = make(map[string]Index)
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesIndexInstruments, exchange, name), nil, nil, &instruments); err != nil {
		return nil, fmt.Errorf("failed to get index instruments: %w", err)
	}
	return instruments, nil
}
