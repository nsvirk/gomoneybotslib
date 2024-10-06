package main

import (
	"log"
	"strconv"
	"strings"

	mbconnect "github.com/nsvirk/gomoneybotslib/pkg/connect"
)

func (t *APITest) InstrumentsInfoBySymbols() {
	symbols := strings.Split(t.cfg.TestSymbols, ",")
	symbolInstruments, err := t.mbClient.InstrumentsInfoBySymbols(symbols)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsBySymbols", symbolInstruments)
}

func (t *APITest) InstrumentsInfoByTokens() {
	tokensStr := strings.Split(t.cfg.TestTokens, ",")
	tokens := make([]uint32, len(tokensStr))
	for i, tokenStr := range tokensStr {
		parsedToken, err := strconv.ParseUint(tokenStr, 10, 32)
		if err != nil {
			log.Fatalf("Error parsing token: %v", err)
		}
		tokens[i] = uint32(parsedToken)
	}
	tokenInstruments, err := t.mbClient.InstrumentsInfoByTokens(tokens)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("InstrumentsByTokens", tokenInstruments)
}

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
	PrettyPrint("InstrumentsByQuery", symbolInstruments)
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
	PrettyPrint("OptionchainTokenSymbolMap", ocTokenSymbolMap)
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
