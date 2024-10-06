package mbconnect

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestInstrumentsQuery(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"NSE:INFY": {
					"instrument_token": 408065,
					"exchange_token": 1594,
					"tradingsymbol": "INFY",
					"name": "INFOSYS LTD",
					"last_price": 1311.65,
					"expiry": "",
					"strike": 0,
					"tick_size": 0.05,
					"lot_size": 1,
					"instrument_type": "EQ",
					"segment": "NSE",
					"exchange": "NSE"
				}
			}
		}`))
	}))
	defer server.Close()

	// Create client with mocked server
	client := New("test_user")
	client.SetBaseURI(server.URL)

	queryParams := InstrumentsQueryParams{
		Exchange:      "NSE",
		Tradingsymbol: "INFY",
	}
	result, err := client.InstrumentsQuery(queryParams)

	if err != nil {
		t.Errorf("InstrumentsByQuery returned an error: %v", err)
	}

	expected := map[string]Instrument{
		"NSE:INFY": {
			InstrumentToken: 408065,
			ExchangeToken:   1594,
			Tradingsymbol:   "INFY",
			Name:            "INFOSYS LTD",
			LastPrice:       1311.65,
			Expiry:          "",
			Strike:          0,
			TickSize:        0.05,
			LotSize:         1,
			InstrumentType:  "EQ",
			Segment:         "NSE",
			Exchange:        "NSE",
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InstrumentsByQuery returned %+v, expected %+v", result, expected)
	}
}

func TestOptionchainInstruments(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"NSE:NIFTY2351116900CE": {
					"instrument_token": 53179655,
					"exchange_token": 207733,
					"tradingsymbol": "NIFTY2351116900CE",
					"name": "NIFTY",
					"last_price": 0,
					"expiry": "2023-05-11",
					"strike": 16900,
					"tick_size": 0.05,
					"lot_size": 50,
					"instrument_type": "CE",
					"segment": "NFO-OPT",
					"exchange": "NFO"
				}
			}
		}`))
	}))
	defer server.Close()

	// Create client with mocked server
	client := New("test_user")
	client.SetBaseURI(server.URL)

	optionChainQueryParams := OptionChainQueryParams{
		Exchange: "NFO",
		Name:     "NIFTY",
	}
	result, err := client.OptionchainInstruments(optionChainQueryParams)

	if err != nil {
		t.Errorf("OptionchainInstruments returned an error: %v", err)
	}

	expected := map[string]Instrument{
		"NSE:NIFTY2351116900CE": {
			InstrumentToken: 53179655,
			ExchangeToken:   207733,
			Tradingsymbol:   "NIFTY2351116900CE",
			Name:            "NIFTY",
			LastPrice:       0,
			Expiry:          "2023-05-11",
			Strike:          16900,
			TickSize:        0.05,
			LotSize:         50,
			InstrumentType:  "CE",
			Segment:         "NFO-OPT",
			Exchange:        "NFO",
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("OptionchainInstruments returned %+v, expected %+v", result, expected)
	}
}

func TestFNOSegmentExpiries(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"NFO-OPT": ["2023-05-11", "2023-05-18", "2023-05-25"],
				"NFO-FUT": ["2023-05-25", "2023-06-29", "2023-07-27"]
			}
		}`))
	}))
	defer server.Close()

	// Create client with mocked server
	client := New("test_user")
	client.SetBaseURI(server.URL)

	result, err := client.FNOSegmentExpiries("NIFTY")

	if err != nil {
		t.Errorf("FNOSegmentExpiries returned an error: %v", err)
	}

	expected := map[string][]string{
		"NFO-OPT": {"2023-05-11", "2023-05-18", "2023-05-25"},
		"NFO-FUT": {"2023-05-25", "2023-06-29", "2023-07-27"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FNOSegmentExpiries returned %+v, expected %+v", result, expected)
	}
}

func TestFNOSegmentNames(t *testing.T) {
	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"data": {
				"NFO-OPT": ["NIFTY", "BANKNIFTY", "FINNIFTY"],
				"NFO-FUT": ["NIFTY", "BANKNIFTY", "FINNIFTY"]
			}
		}`))
	}))
	defer server.Close()

	// Create client with mocked server
	client := New("test_user")
	client.SetBaseURI(server.URL)

	result, err := client.FNOSegmentNames("2023-05-25")

	if err != nil {
		t.Errorf("FNOSegmentNames returned an error: %v", err)
	}

	expected := map[string][]string{
		"NFO-OPT": {"NIFTY", "BANKNIFTY", "FINNIFTY"},
		"NFO-FUT": {"NIFTY", "BANKNIFTY", "FINNIFTY"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FNOSegmentNames returned %+v, expected %+v", result, expected)
	}
}
