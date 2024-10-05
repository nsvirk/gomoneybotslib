package main

import (
	"log"
)

func (t *APITest) IndicesAll() {
	indices, err := t.mbClient.IndicesAll()
	if err != nil {
		log.Fatalf("Error getting indices: %v", err)
	}
	PrettyPrint("IndicesAll", indices)
}

func (t *APITest) IndicesByExchange() {
	indices, err := t.mbClient.IndicesByExchange(t.cfg.TestIndicesExchange)
	if err != nil {
		log.Fatalf("Error getting indices: %v", err)
	}
	PrettyPrint("IndicesByExchange", indices)
}

func (t *APITest) IndexNames() {
	indices, err := t.mbClient.IndexNames(t.cfg.TestIndicesExchange, t.cfg.TestIndicesName)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("IndexNames", indices)
}

func (t *APITest) IndexTokens() {
	tokens, err := t.mbClient.IndexTokens(t.cfg.TestIndicesExchange, t.cfg.TestIndicesName)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("IndexTokens", tokens)
}

func (t *APITest) IndexSymbols() {
	symbols, err := t.mbClient.IndexSymbols(t.cfg.TestIndicesExchange, t.cfg.TestIndicesName)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("IndexSymbols", symbols)
}

func (t *APITest) IndexInstruments() {
	instruments, err := t.mbClient.IndexInstruments(t.cfg.TestIndicesExchange, t.cfg.TestIndicesName)
	if err != nil {
		log.Fatalf("Error getting instruments: %v", err)
	}
	PrettyPrint("IndexInstruments", instruments)
}
