package chunkmap

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverseStr(str string) string {
	bytes := []byte(str)

	i := 0
	j := len(bytes) - 1

	for j > i {
		bytes[i], bytes[j] = bytes[j], bytes[i]

		i += 1
		j -= 1
	}

	return string(bytes)
}

// -----------------------------------------------------------------------

const (
	chunkSize = 64
	zeros64   = "0000000000000000000000000000000000000000000000000000000000000000"
	ones64    = "1111111111111111111111111111111111111111111111111111111111111111"
)

type Bitset struct {
	backend       map[int]uint64
	maxChunkCount int
	maxBitCount   int
	oneBitCount   int
	unknownBitVal uint8
}

func Constructor(size int) Bitset {
	result := Bitset{
		backend:       make(map[int]uint64),
		maxChunkCount: size / chunkSize,
		maxBitCount:   size,
		oneBitCount:   0,
		unknownBitVal: 0,
	}

	if result.maxBitCount%chunkSize != 0 {
		result.maxChunkCount += 1
	}

	return result
}

// -----------------------------------------------------------------------

func (this *Bitset) setBit(idx int, bitVal uint8) {
	chunkIdx := idx / chunkSize
	chunkBitIdx := idx % chunkSize

	chunkVal, ok := this.backend[chunkIdx]
	if !ok {
		if bitVal == this.unknownBitVal {
			return
		}

		chunkVal = 0
		if this.unknownBitVal == 1 {
			chunkVal = math.MaxUint64
		}
	}

	if bitVal == 1 {
		if (chunkVal & (1 << chunkBitIdx)) == 0 {
			this.oneBitCount += 1
		}

		chunkVal |= (uint64(1) << chunkBitIdx)
	} else {
		if (chunkVal & (1 << chunkBitIdx)) != 0 {
			this.oneBitCount -= 1
		}

		chunkVal &= ^(uint64(1) << chunkBitIdx)
	}

	this.backend[chunkIdx] = chunkVal
}

func (this *Bitset) Fix(idx int) {
	if idx < 0 || idx >= this.maxBitCount {
		return
	}

	this.setBit(idx, 1)
}

func (this *Bitset) Unfix(idx int) {
	if idx < 0 || idx >= this.maxBitCount {
		return
	}

	this.setBit(idx, 0)
}

// -----------------------------------------------------------------------

func (this *Bitset) Flip() {
	for k := range this.backend {
		this.backend[k] = ^this.backend[k]
	}

	this.unknownBitVal ^= 1
	this.oneBitCount = this.maxBitCount - this.oneBitCount
}

// -----------------------------------------------------------------------

func (this *Bitset) All() bool {
	return this.Count() == this.maxBitCount
}

func (this *Bitset) One() bool {
	return this.Count() > 0
}

// -----------------------------------------------------------------------

func (this *Bitset) RealCount() int {
	oneCount := 0

	maxChunkCount := this.maxChunkCount
	if this.maxBitCount%chunkSize == 0 {
		maxChunkCount += 1
	}

	for i := 0; i < maxChunkCount-1; i += 1 {
		val, ok := this.backend[i]
		if !ok && this.unknownBitVal == 1 {
			oneCount += chunkSize
			continue
		}

		for val != 0 {
			oneCount += int(val & 1)
			val >>= 1
		}
	}

	lastBitCount := this.maxBitCount % chunkSize

	val, ok := this.backend[this.maxChunkCount-1]
	if !ok && this.unknownBitVal == 1 {
		oneCount += lastBitCount
		return oneCount
	}

	bitPos := 0
	for bitPos < lastBitCount {
		if (val & (1 << bitPos)) != 0 {
			oneCount += 1
		}
		bitPos += 1
	}

	return oneCount
}

func (this *Bitset) Count() int {
	return this.oneBitCount
}

func (this *Bitset) ToString() string {
	if len(this.backend) == 0 {
		buf := make([]byte, this.maxBitCount)
		unknownChunk := zeros64
		if this.unknownBitVal == 1 {
			unknownChunk = ones64
		}
		for pos := 0; pos < this.maxBitCount; {
			n := chunkSize
			if rest := this.maxBitCount - pos; rest < chunkSize {
				n = rest
			}
			copy(buf[pos:pos+n], unknownChunk[:n])
			pos += n
		}
		return string(buf)
	}

	buf := make([]byte, this.maxBitCount)
	unknownChunk := zeros64
	if this.unknownBitVal == 1 {
		unknownChunk = ones64
	}
	pos := 0

	for i := 0; i < this.maxChunkCount; i += 1 {
		restBitCount := this.maxBitCount - pos
		if restBitCount == 0 {
			break
		}

		chunkBitsCount := chunkSize
		if restBitCount < chunkSize {
			chunkBitsCount = restBitCount
		}

		val, ok := this.backend[i]
		if !ok {
			copy(buf[pos:pos+chunkBitsCount], unknownChunk[:chunkBitsCount])
		} else {
			for j := 0; j < chunkBitsCount; j += 1 {
				buf[j+pos] = '0' + byte((val>>j)&1)
			}
		}

		pos += chunkBitsCount
	}

	return string(buf)
}

func (this *Bitset) ToStringWithBuilder() string {
	unknownBitValStr := strconv.Itoa(int(this.unknownBitVal))
	unknownChunkValStr := strings.Repeat(unknownBitValStr, chunkSize)

	var resultStr strings.Builder
	resultStr.Grow(this.maxBitCount)

	usedBitCount := 0

	for i := 0; i < this.maxChunkCount; i += 1 {
		val, ok := this.backend[i]
		if !ok {
			restBitCount := this.maxBitCount - usedBitCount
			if restBitCount >= chunkSize {
				resultStr.WriteString(unknownChunkValStr)
				usedBitCount += chunkSize
				continue
			}

			resultStr.WriteString(strings.Repeat(unknownBitValStr, restBitCount))
			usedBitCount += restBitCount
			break
		}

		if i+1 == this.maxChunkCount {
			partResultStr := ""
			restBitCount := this.maxBitCount - usedBitCount
			for restBitCount > 0 {
				partResultStr += strconv.Itoa(int(val & 1))
				val >>= 1
				restBitCount -= 1
			}
			usedBitCount += restBitCount
			resultStr.WriteString(partResultStr)
			break
		}

		resultStr.WriteString(reverseStr(fmt.Sprintf("%064b", val)))
		usedBitCount += chunkSize
	}

	return resultStr.String()
}
