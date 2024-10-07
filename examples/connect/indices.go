package main

import (
	"fmt"
	"log"
)

// -----------------------------------------------------
// /indices/all
// -----------------------------------------------------
func (t *APITest) IndicesAll() {
	indices, err := t.mbClient.IndicesAll()
	if err != nil {
		log.Fatalf("Error getting indices: %v", err)
	}
	PrettyPrint("IndicesAll", "", uint32(len(indices)), indices)
}

// -----------------------------------------------------
// /indices/:exchange
// -----------------------------------------------------
func (t *APITest) IndicesByExchange() {
	testIndicesExchange := t.cfg.TestIndicesExchange
	// testIndicesExchange = ""
	indices, err := t.mbClient.IndicesByExchange(testIndicesExchange)
	if err != nil {
		log.Fatalf("Error getting indices: %v", err)
	}
	PrettyPrint("IndicesByExchange", testIndicesExchange, uint32(len(indices)), indices)
}

// -----------------------------------------------------
// /indices/:exchange/:name
// -----------------------------------------------------
func (t *APITest) IndexInstruments() {
	testIndicesExchange := t.cfg.TestIndicesExchange
	testIndicesName := t.cfg.TestIndicesName
	// testIndicesExchange = ""
	// testIndicesName = ""
	instruments, err := t.mbClient.IndexInstruments(testIndicesExchange, testIndicesName)
	if err != nil {
		log.Fatalf("Error getting index instruments: %v", err)
	}
	PrettyPrint("IndexInstruments", fmt.Sprintf("%s %s", testIndicesExchange, testIndicesName), uint32(len(instruments)), instruments)
}
