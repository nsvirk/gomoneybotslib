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

// -----------------------------------------------------
// /indices/all
// -----------------------------------------------------
// IndicesAll returns all the indices
func (c *Client) IndicesAll() (map[string][]Index, error) {
	var index = make(map[string][]Index)
	if err := c.doEnvelope(http.MethodGet, URIIndicesAll, nil, nil, &index); err != nil {
		return nil, err
	}
	return index, nil
}

// -----------------------------------------------------
// /indices/:exchange
// -----------------------------------------------------
// IndicesByExchange returns the indices for a given exchange
func (c *Client) IndicesByExchange(exchange string) ([]Index, error) {
	if exchange == "" {
		return nil, fmt.Errorf("`exchange` is required")
	}
	var indices []Index
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesByExchange, exchange), nil, nil, &indices); err != nil {
		return nil, err
	}
	return indices, nil
}

// -----------------------------------------------------
// /indices/:exchange/:name
// -----------------------------------------------------
// IndexInstruments returns the instruments of the indices for a Index name
func (c *Client) IndexInstruments(exchange, name string) ([]Index, error) {
	if exchange == "" {
		return nil, fmt.Errorf("`exchange` is required")
	}
	if name == "" {
		return nil, fmt.Errorf("`name` is required")
	}
	var indices []Index
	if err := c.doEnvelope(http.MethodGet, fmt.Sprintf(URIIndicesIndexInstruments, exchange, name), nil, nil, &indices); err != nil {
		return nil, err
	}
	return indices, nil
}
