package mbconnect

import (
	"fmt"
	"net/http"
	"time"
)

// Index is a struct that represents an index
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

// GET /indices/all - Get all indices info `exchange` wise
func (c *Client) IndicesAll() (map[string][]Index, error) {
	var index = make(map[string][]Index)
	if err := c.doEnvelope(http.MethodGet, URIIndicesAll, nil, nil, &index); err != nil {
		return nil, err
	}
	return index, nil
}

// GET /indices/:exchange/info - Get indices info by `exchange`
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

// GET /indices/:exchange/:name/instruments - Get instruments of the indices by `exchange` and `name`
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
