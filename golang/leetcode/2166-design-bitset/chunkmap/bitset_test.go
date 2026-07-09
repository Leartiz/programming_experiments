package chunkmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Bitset_ToString(t *testing.T) {
	bitset := Constructor(2)
	got := bitset.ToString()
	want := "00"
	t.Logf("got: %v, want: %v", got, want)
	assert.Equal(t, want, got)

	bitset.Flip()
	got = bitset.ToString()
	want = "11"
	t.Logf("got: %v, want: %v", got, want)
	assert.Equal(t, want, got)

	bitset.Unfix(0)
	got = bitset.ToString()
	want = "01"
	t.Logf("got: %v, want: %v", got, want)
	assert.Equal(t, want, got)
}

func Test_Bitset_One(t *testing.T) {
	bitset := Constructor(2)
	got := bitset.One()
	want := false
	assert.Equal(t, want, got)
}

func Test_Bitset_All(t *testing.T) {
	bitset := Constructor(2)
	got := bitset.All()
	want := false
	assert.Equal(t, want, got)
}

func Test_Bitset_All_1(t *testing.T) {
	bitset := Constructor(2)
	bitset.Fix(0)
	bitset.Fix(1)
	got := bitset.All()
	want := true
	assert.Equal(t, want, got)
}

func Test_Bitset_Count(t *testing.T) {
	bitset := Constructor(2)
	bitset.Fix(0)
	bitset.Fix(1)
	got := bitset.Count()
	want := 2
	assert.Equal(t, want, got)
}

func Test_Bitset_Bitset_90(t *testing.T) {
	bitset := Constructor(90)
	bitset.Fix(0)
	bitset.Fix(64)
	got := bitset.ToString()
	want := "100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000"
	assert.Equal(t, want, got)
}

func Test_Bitset_Bitset_64(t *testing.T) {
	bitset := Constructor(64)
	bitset.Fix(53)
	got := bitset.Count()
	want := 1
	assert.Equal(t, want, got)
}
