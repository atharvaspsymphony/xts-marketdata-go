package marketdata

import (
	"bytes"
	"compress/flate"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

// ApplicationMessageVersion enum
type ApplicationMessageVersion int

const (
	Version_1           ApplicationMessageVersion = 1
	Version_1_0_1_0969  ApplicationMessageVersion = 2
	Version_1_0_1_2879  ApplicationMessageVersion = 3
	Version_1_0_1_2983  ApplicationMessageVersion = 4
)

// BinaryReader struct
type BinaryReader struct {
	data   []byte
	offset int
}

func NewBinaryReader(data []byte) *BinaryReader {
	return &BinaryReader{data: data, offset: 0}
}

func (br *BinaryReader) ReadInt8() int8 {
	val := int8(br.data[br.offset])
	br.offset++
	return val
}

func (br *BinaryReader) ReadUint8() uint8 {
	val := br.data[br.offset]
	br.offset++
	return val
}

func (br *BinaryReader) ReadInt16() int16 {
	val := int16(binary.LittleEndian.Uint16(br.data[br.offset:]))
	br.offset += 2
	return val
}

func (br *BinaryReader) ReadUint16() uint16 {
	val := binary.LittleEndian.Uint16(br.data[br.offset:])
	br.offset += 2
	return val
}

func (br *BinaryReader) ReadInt32() int32 {
	val := int32(binary.LittleEndian.Uint32(br.data[br.offset:]))
	br.offset += 4
	return val
}

func (br *BinaryReader) ReadUint32() uint32 {
	val := binary.LittleEndian.Uint32(br.data[br.offset:])
	br.offset += 4
	return val
}

func (br *BinaryReader) ReadInt64() int64 {
	val := int64(binary.LittleEndian.Uint64(br.data[br.offset:]))
	br.offset += 8
	return val
}

func (br *BinaryReader) ReadUint64() uint64 {
	val := binary.LittleEndian.Uint64(br.data[br.offset:])
	br.offset += 8
	return val
}

func (br *BinaryReader) ReadFloat64() float64 {
	bits := binary.LittleEndian.Uint64(br.data[br.offset:])
	br.offset += 8
	return math.Float64frombits(bits)
}

func (br *BinaryReader) ReadBytes(n int) []byte {
	val := br.data[br.offset : br.offset+n]
	br.offset += n
	return val
}

// MarketDeptRowInfo struct
type MarketDeptRowInfo struct {
	Size                int64   `json:"size"`
	RowPrice            float64 `json:"rowprice"`
	TotalOrders         uint32  `json:"totalOrders"`
	BackMarketMakerFlag int16   `json:"backmarketmakerflag"`
}

func DeserializeMarketDeptRowInfo(reader *BinaryReader) (MarketDeptRowInfo, int) {
	count := 0
	row := MarketDeptRowInfo{}

	row.Size = reader.ReadInt64()
	count += 8

	row.RowPrice = reader.ReadFloat64()
	count += 8

	row.TotalOrders = reader.ReadUint32()
	count += 4

	row.BackMarketMakerFlag = reader.ReadInt16()
	count += 2

	return row, count
}

// MarketDepthEvent struct
type MarketDepthEventWS struct {
	MessageVersion       uint16              `json:"messageVersion"`
	ApplicationType      uint16              `json:"applicationType"`
	TokenID              uint64              `json:"tokenID"`
	SequenceNumber       uint64              `json:"sequenceNumber"`
	SkipBytes            int32               `json:"skipBytes"`
	ExchangeSegment      int16               `json:"exchangeSegment"`
	ExchangeInstrumentID int32               `json:"exchangeInstrumentId"`
	ExchangeTimestamp    uint64              `json:"exchangeTimestamp"`
	BidCount             int32               `json:"bidCount"`
	Bid                  []MarketDeptRowInfo `json:"Bid"`
	AskCount             int32               `json:"askCount"`
	Ask                  []MarketDeptRowInfo `json:"Ask"`
	LUT                  uint64              `json:"lut"`
	LTP                  float64             `json:"LTP"`
	LTQ                  int64               `json:"ltq"`
	TotalBuyQuantity     int64               `json:"totalBuyQuantity"`
	TotalSellQuantity    int64               `json:"totalSellQuantity"`
	TotalTradedQuantity  int64               `json:"totalTradedQuantity"`
	AverageTradedPrice   float64             `json:"averageTradedPrice"`
	LastTradedTime       int64               `json:"lastTradedTime"`
	PercentChange        float64             `json:"percentChange"`
	Open                 float64             `json:"open"`
	High                 float64             `json:"high"`
	Low                  float64             `json:"low"`
	Close                float64             `json:"close"`
	TotalValueTraded     float64             `json:"totalValueTraded"`
	BBTotalBuy           int16               `json:"bbTotalBuy"`
	BBTotalSell          int16               `json:"bbTotalSell"`
	BookType             int16               `json:"BookType"`
	MarketType           int16               `json:"MarketType"`
}

func DeserializeMarketDepthEvent(reader *BinaryReader, count int) map[string]interface{} {
	count += 2
	messageVersion := reader.ReadUint16()
	applicationType := reader.ReadUint16()
	tokenID := reader.ReadUint64()
	count += 8

	var sequenceNumber uint64
	var skipBytes int32

	if messageVersion >= uint16(Version_1_0_1_2983) {
		sequenceNumber = reader.ReadUint64()
		count += 8
		skipBytes = reader.ReadInt32()
		count += 4
	}

	exchangeSegment := reader.ReadInt16()
	count += 2

	exchangeInstrumentID := reader.ReadInt32()
	count += 4

	exchangeTimestamp := reader.ReadUint64()
	count += 8

	bidCount := reader.ReadInt32()
	count += 4

	bidData := make([]MarketDeptRowInfo, bidCount)
	for i := 0; i < int(bidCount); i++ {
		row, rowCount := DeserializeMarketDeptRowInfo(reader)
		count += rowCount
		bidData[i] = row
	}

	askCount := reader.ReadInt32()
	count += 4

	askData := make([]MarketDeptRowInfo, askCount)
	for i := 0; i < int(askCount); i++ {
		row, rowCount := DeserializeMarketDeptRowInfo(reader)
		count += rowCount
		askData[i] = row
	}

	// Skip two MarketDeptRowInfo entries
	_, rowCount1 := DeserializeMarketDeptRowInfo(reader)
	count += rowCount1
	_, rowCount2 := DeserializeMarketDeptRowInfo(reader)
	count += rowCount2

	lut := reader.ReadUint64()
	count += 8

	ltp := reader.ReadFloat64()
	count += 8

	ltq := reader.ReadInt64()
	count += 8

	totalBuyQuantity := reader.ReadInt64()
	count += 8

	totalSellQuantity := reader.ReadInt64()
	count += 8

	totalTradedQuantity := reader.ReadInt64()
	count += 8

	averageTradedPrice := reader.ReadFloat64()
	count += 8

	lastTradedTime := reader.ReadInt64()
	count += 8

	percentChange := reader.ReadFloat64()
	count += 8

	open := reader.ReadFloat64()
	count += 8

	high := reader.ReadFloat64()
	count += 8

	low := reader.ReadFloat64()
	count += 8

	close := reader.ReadFloat64()
	count += 8

	totalValueTraded := reader.ReadFloat64()
	totalValueTraded = 0
	count += 8

	bbTotalBuy := reader.ReadInt16()
	count += 2

	bbTotalSell := reader.ReadInt16()
	count += 2

	bookType := reader.ReadInt16()
	count += 2

	marketType := reader.ReadInt16()

	return map[string]interface{}{
		"Marketdepth": map[string]interface{}{
			"messageVersion":       messageVersion,
			"applicationType":      applicationType,
			"tokenID":              tokenID,
			"sequenceNumber":       sequenceNumber,
			"skipBytes":            skipBytes,
			"exchangeSegment":      exchangeSegment,
			"exchangeInstrumentId": exchangeInstrumentID,
			"exchangeTimestamp":    exchangeTimestamp,
			"bidCount":             bidCount,
			"Bid":                  bidData,
			"askCount":             askCount,
			"Ask":                  askData,
			"lut":                  lut,
			"LTP":                  ltp,
			"ltq":                  ltq,
			"totalBuyQuantity":     totalBuyQuantity,
			"totalSellQuantity":    totalSellQuantity,
			"totalTradedQuantity":  totalTradedQuantity,
			"averageTradedPrice":   averageTradedPrice,
			"lastTradedTime":       lastTradedTime,
			"percentChange":        percentChange,
			"open":                 open,
			"high":                 high,
			"low":                  low,
			"close":                close,
			"totalValueTraded":     totalValueTraded,
			"bbTotalBuy":           bbTotalBuy,
			"bbTotalSell":          bbTotalSell,
			"BookType":             bookType,
			"MarketType":           marketType,
		},
	}
}

// OpenInterest struct
type OpenInterestWS struct {
	ExchangeSegment             int16  `json:"exchangeSegment"`
	ExchangeInstrumentID        int32  `json:"exchangeInstrumentId"`
	ExchangeTimestamp           uint64 `json:"exchangeTimestamp"`
	OpenInterest                int64  `json:"openInterest"`
	UnderlyingExchangeSegment   int16  `json:"underlyingExchangeSegment"`
	UnderlyingInstrumentID      uint64 `json:"underlyingInstrumentID"`
	IsStringExits               int8   `json:"isStringExits"`
	underlyingIDIndexName	    string `json:"underlyingIDIndexName"`
	UnderlyingTotalOpenInterest uint64 `json:"underlyingTotalOpenInterest"`
}

func DeserializeOpenInterest(reader *BinaryReader, count int) map[string]interface{} {
	count += 2
	messageVersion := reader.ReadUint16()
	_ = reader.ReadUint16()
	_ = reader.ReadUint64()
	count += 8

	if messageVersion >= uint16(Version_1_0_1_2983) {
		_ = reader.ReadUint64()
		count += 8
		_ = reader.ReadInt32()
		count += 4
	}

	exchangeSegment := reader.ReadInt16()
	count += 2

	exchangeInstrumentID := reader.ReadInt32()
	count += 4

	exchangeTimestamp := reader.ReadUint64()
	count += 8

	_ = reader.ReadInt16()
	count += 2
	// fmt.Printf("marketType--> %d\n", marketType)

	openInterest := reader.ReadInt64()
	count += 8

	underlyingExchangeSegment := reader.ReadInt16()
	count += 2

	underlyingInstrumentID := reader.ReadUint64()
	count += 8

	isStringExits := reader.ReadInt8()
	count += 1

	var underlyingIDIndexName string

	if isStringExits == 1 {
		stringLength := reader.ReadInt8()
		count++

		data := reader.ReadBytes(int(stringLength)) // read bytes
		count += int(stringLength)

		underlyingIDIndexName = strings.ToUpper(string(data))
		count += int(stringLength)
	}

	underlyingTotalOpenInterest := reader.ReadUint64()
	count += 8

	return map[string]interface{}{
		"OpenInterest": map[string]interface{}{
			"exchangeSegment":             exchangeSegment,
			"exchangeInstrumentId":        exchangeInstrumentID,
			"exchangeTimestamp":           exchangeTimestamp,
			"openInterest":                openInterest,
			"underlyingExchangeSegment":   underlyingExchangeSegment,
			"underlyingInstrumentID":      underlyingInstrumentID,
			"isStringExits":               isStringExits,
			"underlyingIDIndexName":       underlyingIDIndexName,
			"underlyingTotalOpenInterest": underlyingTotalOpenInterest,
		},
	}
}

// Touchline struct
type TouchlineWS struct {
	MessageVersion       uint16            `json:"messageVersion"`
	ApplicationType      uint16            `json:"applicationType"`
	TokenID              uint64            `json:"tokenID"`
	SequenceNumber       uint64            `json:"sequenceNumber"`
	SkipBytes            int32             `json:"skipBytes"`
	ExchangeSegment      int16             `json:"exchangeSegment"`
	ExchangeInstrumentID int32             `json:"exchangeInstrumentId"`
	ExchangeTimestamp    uint64            `json:"exchangeTimestamp"`
	Bid                  MarketDeptRowInfo `json:"Bid"`
	Ask                  MarketDeptRowInfo `json:"Ask"`
	LUT                  uint64            `json:"lut"`
	LTP                  float64           `json:"LTP"`
	LTQ                  int64             `json:"ltq"`
	TotalBuyQuantity     int64             `json:"totalBuyQuantity"`
	TotalSellQuantity    int64             `json:"totalSellQuantity"`
	TotalTradedQuantity  int64             `json:"totalTradedQuantity"`
	AverageTradedPrice   float64           `json:"averageTradedPrice"`
	LastTradedTime       int64             `json:"lastTradedTime"`
	PercentChange        float64           `json:"percentChange"`
	Open                 float64           `json:"open"`
	High                 float64           `json:"high"`
	Low                  float64           `json:"low"`
	Close                float64           `json:"close"`
	TotalValueTraded     float64           `json:"totalValueTraded"`
	BBTotalBuy           int16             `json:"bbTotalBuy"`
	BBTotalSell          int16             `json:"bbTotalSell"`
	BookType             int16             `json:"BookType"`
	MarketType           int16             `json:"MarketType"`
}

func DeserializeTouchline(reader *BinaryReader, count int) map[string]interface{} {
	count += 2
	messageVersion := reader.ReadUint16()
	applicationType := reader.ReadUint16()
	tokenID := reader.ReadUint64()
	// fmt.Printf("messageVersion--> %d applicationType--> %d tokenID--> %d\n", messageVersion, applicationType, tokenID)
	count += 8

	var sequenceNumber uint64
	var skipBytes int32

	if messageVersion >= uint16(Version_1_0_1_2983) {
		sequenceNumber = reader.ReadUint64()
		count += 8
		skipBytes = reader.ReadInt32()
		count += 4
	}

	exchangeSegment := reader.ReadInt16()
	count += 2

	exchangeInstrumentID := reader.ReadInt32()
	count += 4

	exchangeTimestamp := reader.ReadUint64()
	count += 8

	bidData, bidCount := DeserializeMarketDeptRowInfo(reader)
	count += bidCount

	askData, askCount := DeserializeMarketDeptRowInfo(reader)
	count += askCount

	lut := reader.ReadUint64()
	count += 8

	ltp := reader.ReadFloat64()
	count += 8

	ltq := reader.ReadInt64()
	count += 8

	totalBuyQuantity := reader.ReadInt64()
	count += 8

	totalSellQuantity := reader.ReadInt64()
	count += 8

	totalTradedQuantity := reader.ReadInt64()
	count += 8

	averageTradedPrice := reader.ReadFloat64()
	count += 8

	lastTradedTime := reader.ReadInt64()
	count += 8

	percentChange := reader.ReadFloat64()
	count += 8

	open := reader.ReadFloat64()
	count += 8

	high := reader.ReadFloat64()
	count += 8

	low := reader.ReadFloat64()
	count += 8

	close := reader.ReadFloat64()
	count += 8

	totalValueTraded := reader.ReadFloat64()
	totalValueTraded = 0
	count += 8

	bbTotalBuy := reader.ReadInt16()
	count += 2

	bbTotalSell := reader.ReadInt16()
	count += 2

	bookType := reader.ReadInt16()
	count += 2

	marketType := reader.ReadInt16()

	return map[string]interface{}{
		"Touchline": map[string]interface{}{
			"messageVersion":       messageVersion,
			"applicationType":      applicationType,
			"tokenID":              tokenID,
			"sequenceNumber":       sequenceNumber,
			"skipBytes":            skipBytes,
			"exchangeSegment":      exchangeSegment,
			"exchangeInstrumentId": exchangeInstrumentID,
			"exchangeTimestamp":    exchangeTimestamp,
			"Bid":                  bidData,
			"Ask":                  askData,
			"lut":                  lut,
			"LTP":                  ltp,
			"ltq":                  ltq,
			"totalBuyQuantity":     totalBuyQuantity,
			"totalSellQuantity":    totalSellQuantity,
			"totalTradedQuantity":  totalTradedQuantity,
			"averageTradedPrice":   averageTradedPrice,
			"lastTradedTime":       lastTradedTime,
			"percentChange":        percentChange,
			"open":                 open,
			"high":                 high,
			"low":                  low,
			"close":                close,
			"totalValueTraded":     totalValueTraded,
			"bbTotalBuy":           bbTotalBuy,
			"bbTotalSell":          bbTotalSell,
			"BookType":             bookType,
			"MarketType":           marketType,
		},
	}
}

// pakoInflateRaw decompresses raw DEFLATE data (without zlib wrapper)
func pakoInflateRaw(data []byte) ([]byte, error) {
	reader := flate.NewReader(bytes.NewReader(data))
	defer reader.Close()

	var out bytes.Buffer
	_, err := io.Copy(&out, reader)
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func onXTsBinaryPacket(data []byte) ([]map[string]interface{}, error) {
	offset := 0
	count := 0
	dataLen := len(data)
	isNextPacket := true
	var results []map[string]interface{}

	for isNextPacket {
		nextData := data[offset:dataLen]
		br := NewBinaryReader(nextData)

		isGzipCompressed := br.ReadInt8()
		offset++

		if isGzipCompressed == 1 {
			nextData = data[offset:dataLen]
			br = NewBinaryReader(nextData)

			_ = br.ReadUint16() // messageCode
			_ = br.ReadInt16()  // exchangeSegment
			_ = br.ReadInt32()  // exchangeInstrumentID
			_ = br.ReadInt16()  // bookType
			_ = br.ReadInt16()  // marketType
			_ = br.ReadUint16()
			compressedPacketSize := br.ReadUint16()
			offset += 16

			filteredByteArray := data[offset : offset+int(compressedPacketSize)]
			inflate, err := pakoInflateRaw(filteredByteArray)
			if err != nil {
				return nil, fmt.Errorf("error decompressing: %v", err)
			}

			r := NewBinaryReader(inflate)
			currentSize := int(compressedPacketSize) + offset
			if currentSize < len(data) {
				isNextPacket = true
				offset = currentSize
			} else {
				isNextPacket = false
			}

			messageCodeStr := r.ReadUint16()

			var jsonData map[string]interface{}
			switch messageCodeStr {
			case 1501:
				jsonData = DeserializeTouchline(r, count)
			case 1502:
				jsonData = DeserializeMarketDepthEvent(r, count)
			case 1510:
				jsonData = DeserializeOpenInterest(r, count)
			default:
				continue
			}

			results = append(results, jsonData)

		} else if isGzipCompressed == 0 {
			messageCode := br.ReadUint16()
			_ = br.ReadInt16()  // exchangeSegment
			_ = br.ReadInt32()  // exchangeInstrumentID
			_ = br.ReadInt16()  // bookType
			_ = br.ReadInt16()  // marketType
			uncompressedPacketSize := br.ReadUint16()
			_ = br.ReadUint16() // compressedPacketSize
			offset += 14
			count = offset

			var jsonData map[string]interface{}
			switch messageCode {
			case 1501:
				jsonData = DeserializeTouchline(br, count)
			case 1502:
				jsonData = DeserializeMarketDepthEvent(br, count)
			case 1510:
				jsonData = DeserializeOpenInterest(br, count)
			default:
				continue
			}

			results = append(results, jsonData)

			currentSize := offset + int(uncompressedPacketSize)
			if currentSize < len(data) {
				isNextPacket = true
				offset = currentSize
			} else {
				isNextPacket = false
			}
		}
	}

	return results, nil
}
