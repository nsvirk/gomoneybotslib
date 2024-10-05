package main

import (
	"log"
	"strconv"
	"strings"

	mbconnect "github.com/nsvirk/gomoneybotslib/pkg/connect"
)

func (t *APITest) InstrumentsBySymbols() {
	symbols := strings.Split(t.cfg.TestSymbols, ",")
	symbolInstruments, err := t.mbClient.InstrumentsBySymbols(symbols)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsBySymbols", symbolInstruments)
}

func (t *APITest) InstrumentsByTokens() {
	tokensStr := strings.Split(t.cfg.TestTokens, ",")
	tokens := make([]uint32, len(tokensStr))
	for i, tokenStr := range tokensStr {
		parsedToken, err := strconv.ParseUint(tokenStr, 10, 32)
		if err != nil {
			log.Fatalf("Error parsing token: %v", err)
		}
		tokens[i] = uint32(parsedToken)
	}
	tokenInstruments, err := t.mbClient.InstrumentsByTokens(tokens)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsByTokens", tokenInstruments)
}

func (t *APITest) InstrumentsByQuery() {
	qp := mbconnect.InstrumentsQueryParams{
		Exchange:       "NFO",
		Name:           "NIFTY",
		InstrumentType: "FUT",
	}
	symbolInstruments, err := t.mbClient.InstrumentsByQuery(qp)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsByQuery", symbolInstruments)
}

func (t *APITest) SymbolsByQuery() {
	qp := mbconnect.InstrumentsQueryParams{
		Exchange:       t.cfg.TestQueryExchange,
		Name:           t.cfg.TestQueryName,
		InstrumentType: t.cfg.TestQueryInstrumentType,
	}
	symbols, err := t.mbClient.SymbolsByQuery(qp)
	if err != nil {
		log.Fatalf("Error getting instrument symbols: %v", err)
	}
	PrettyPrint("SymbolsByQuery", symbols)
}

func (t *APITest) TokensByQuery() {
	qp := mbconnect.InstrumentsQueryParams{
		Exchange:       t.cfg.TestQueryExchange,
		Name:           t.cfg.TestQueryName,
		InstrumentType: t.cfg.TestQueryInstrumentType,
	}
	tokens, err := t.mbClient.TokensByQuery(qp)
	if err != nil {
		log.Fatalf("Error getting instrument tokens: %v", err)
	}
	PrettyPrint("TokensByQuery", tokens)
}

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
	PrettyPrint("OptionchainInstruments", ocInstruments)
}

func (t *APITest) OptionchainSymbols() {
	ocp := mbconnect.OptionChainQueryParams{
		Exchange:  t.cfg.TestOCExchange,
		Name:      t.cfg.TestOCName,
		FutExpiry: t.cfg.TestOCFutExpiry,
		OptExpiry: t.cfg.TestOCOptExpiry,
	}
	ocSymbols, err := t.mbClient.OptionchainSymbols(ocp)
	if err != nil {
		log.Fatalf("Error getting option chain symbols: %v", err)
	}
	PrettyPrint("OptionchainSymbols", ocSymbols)
}

func (t *APITest) OptionchainTokens() {
	ocp := mbconnect.OptionChainQueryParams{
		Exchange:  t.cfg.TestOCExchange,
		Name:      t.cfg.TestOCName,
		FutExpiry: t.cfg.TestOCFutExpiry,
		OptExpiry: t.cfg.TestOCOptExpiry,
	}
	ocTokens, err := t.mbClient.OptionchainTokens(ocp)
	if err != nil {
		log.Fatalf("Error getting option chain tokens: %v", err)
	}
	PrettyPrint("OptionchainTokens", ocTokens)
}

func (t *APITest) FNOSegmentExpiries() {
	segmentExpiry, err := t.mbClient.FNOSegmentExpiries(t.cfg.TestSegmentName)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("FNOSegmentExpiries", segmentExpiry)
}

func (t *APITest) FNOSegmentNames() {
	segmentName, err := t.mbClient.FNOSegmentNames(t.cfg.TestSegmentExpiry)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("FNOSegmentNames", segmentName)
}
