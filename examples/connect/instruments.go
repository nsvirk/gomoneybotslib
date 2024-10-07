package main

import (
	"log"
	"strconv"
	"strings"

	mbconnect "github.com/nsvirk/gomoneybotslib/pkg/connect"
)

// -----------------------------------------------------
// /instruments/info?s=NSE:SBIN&s=BSE:SBIN
// -----------------------------------------------------
func (t *APITest) InstrumentsInfoBySymbols() {
	testSymbols := t.cfg.TestSymbols
	// var symbols []string
	symbols := strings.Split(testSymbols, ",")
	symbolInstruments, err := t.mbClient.InstrumentsInfoBySymbols(symbols)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsBySymbols", testSymbols, uint32(len(symbolInstruments)), symbolInstruments)
}

// -----------------------------------------------------
// /instruments/info?t=256265&t=8961794
// -----------------------------------------------------
func (t *APITest) InstrumentsInfoByTokens() {
	testTokensStr := t.cfg.TestTokens
	tokens, err := stringToTokens(testTokensStr)
	if err != nil {
		log.Fatalf("Error formatting tokens: %v", err)
	}
	// var tokens []uint32
	tokenInstruments, err := t.mbClient.InstrumentsInfoByTokens(tokens)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsByTokens", testTokensStr, uint32(len(tokenInstruments)), tokenInstruments)
}

// -----------------------------------------------------
// /instruments/query
// -----------------------------------------------------
func (t *APITest) InstrumentsQuery() {
	qp := mbconnect.InstrumentsQueryParams{
		Exchange:       "NFO",
		Name:           "NIFTY",
		InstrumentType: "FUT",
	}
	symbolInstruments, err := t.mbClient.InstrumentsQuery(qp)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsByQuery", qp, uint32(len(symbolInstruments)), symbolInstruments)
}

// -----------------------------------------------------
// /instruments/fno/optionchain
// -----------------------------------------------------
func (t *APITest) OptionchainInstruments() {
	ocp := mbconnect.OptionChainQueryParams{
		Exchange:  t.cfg.TestOCExchange,
		Name:      t.cfg.TestOCName,
		FutExpiry: t.cfg.TestOCFutExpiry,
		OptExpiry: t.cfg.TestOCOptExpiry,
	}
	ocInstruments, err := t.mbClient.OptionchainInstruments(ocp)
	if err != nil {
		log.Fatalf("Error getting option chain instruments: %v", err)
	}
	PrettyPrint("OptionchainInstruments", ocp, uint32(len(ocInstruments)), ocInstruments)
}

func (t *APITest) OptionchainTokenSymbolMap() {
	ocp := mbconnect.OptionChainQueryParams{
		Exchange:  t.cfg.TestOCExchange,
		Name:      t.cfg.TestOCName,
		FutExpiry: t.cfg.TestOCFutExpiry,
		OptExpiry: t.cfg.TestOCOptExpiry,
	}
	ocTokenSymbolMap, err := t.mbClient.OptionchainTokenSymbolMap(ocp)
	if err != nil {
		log.Fatalf("Error getting option chain token symbol map: %v", err)
	}
	PrettyPrint("OptionchainTokenSymbolMap", ocp, uint32(len(ocTokenSymbolMap)), ocTokenSymbolMap)
}

// -----------------------------------------------------
// /instruments/fno/segment_expiries/:name
// -----------------------------------------------------
func (t *APITest) FNOSegmentExpiries() {
	testSegmentName := t.cfg.TestSegmentName
	// testSegmentName = ""
	segmentExpiry, err := t.mbClient.FNOSegmentExpiries(testSegmentName)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("FNOSegmentExpiries", testSegmentName, uint32(len(segmentExpiry)), segmentExpiry)
}

// -----------------------------------------------------
// /instruments/fno/segment_names/:expiry
// -----------------------------------------------------
func (t *APITest) FNOSegmentNames() {
	testSegmentExpiry := t.cfg.TestSegmentExpiry
	// testSegmentExpiry = ""
	segmentName, err := t.mbClient.FNOSegmentNames(testSegmentExpiry)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("FNOSegmentNames", testSegmentExpiry, uint32(len(segmentName)), segmentName)
}

func stringToTokens(testTokensStr string) ([]uint32, error) {
	tokensStr := strings.Split(testTokensStr, ",")
	tokens := make([]uint32, len(tokensStr))
	for i, tokenStr := range tokensStr {
		parsedToken, err := strconv.ParseUint(tokenStr, 10, 32)
		if err != nil {
			return nil, err
		}
		tokens[i] = uint32(parsedToken)
	}
	return tokens, nil
}
