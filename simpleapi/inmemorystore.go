package simpleapi

import (
	"sync"
)

var (
	list = make(map[string]interface{})
	mu   sync.RWMutex
)

// LoadInMemory stores market data in memory.
func LoadInMemory(messageCode, exchangeSegment, exchangeInstrumentID string, marketData interface{}) interface{} {
	key := messageCode + "|" + exchangeSegment + "|" + exchangeInstrumentID
	mu.Lock()
	defer mu.Unlock()
	list[key] = marketData
	return marketData
}

// GetFromInMemory retrieves market data from memory.
func GetFromInMemory(messageCode, exchangeSegment, exchangeInstrumentID string) interface{} {
	key := messageCode + "|" + exchangeSegment + "|" + exchangeInstrumentID
	mu.RLock()
	defer mu.RUnlock()
	return list[key]
}
