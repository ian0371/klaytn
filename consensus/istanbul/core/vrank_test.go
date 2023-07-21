package core

import (
	"testing"
	"time"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/consensus/istanbul"
	"github.com/klaytn/klaytn/consensus/istanbul/validator"
	"github.com/stretchr/testify/assert"
)

func TestVrankCategorizeArrivalTimeMap(t *testing.T) {
	addrs, _ := genValidators(6)
	committee := []istanbul.Validator{}
	src := make(map[common.Address]time.Duration)
	for i, addr := range addrs {
		committee = append(committee, validator.New(addr))
		src[addr] = time.Duration((i * 100) * int(time.Millisecond))
	}
	categorized := categorizeArrivalTimeMap(committee, src)
	var cntEarly, cntLate int
	for _, addr := range addrs {
		if categorized[addr] == vrankArrivedEarly {
			cntEarly++
		} else if categorized[addr] == vrankArrivedLate {
			cntLate++
		}
	}
	// threshold: vrankDefaultLateThreshold (300ms)
	assert.Equal(t, 4, cntEarly)
	assert.Equal(t, 2, cntLate)
}

func TestCompressSerializedArrivals(t *testing.T) {
	tcs := []struct {
		input    []int
		expected []byte
	}{
		{
			input:    []int{2},
			expected: []byte{0b10_000000},
		},
		{
			input:    []int{2, 1},
			expected: []byte{0b10_01_0000},
		},
		{
			input:    []int{0, 2, 1},
			expected: []byte{0b00_10_01_00},
		},
		{
			input:    []int{0, 2, 1, 1},
			expected: []byte{0b00_10_01_01},
		},
		{
			input:    []int{1, 2, 1, 2, 1},
			expected: []byte{0b01_10_01_10, 0b01_000000},
		},
		{
			input:    []int{1, 2, 1, 2, 1, 2},
			expected: []byte{0b01_10_01_10, 0b01_10_0000},
		},
		{
			input:    []int{1, 2, 1, 2, 1, 2, 1},
			expected: []byte{0b01_10_01_10, 0b01_10_01_00},
		},
		{
			input:    []int{1, 2, 1, 2, 1, 2, 0, 2},
			expected: []byte{0b01_10_01_10, 0b01_10_00_10},
		},
		{
			input:    []int{1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 1, 2, 2, 1, 1},
			expected: []byte{0b01011010, 0b01011010, 0b01011010, 0b01011010, 0b01010000},
		},
	}
	for i, tc := range tcs {
		assert.Equal(t, compressSerializedArrivals(tc.input), tc.expected, "tc %d failed", i)
	}
}
