package dummyfile

import (
	"bufio"
	"bytes"
	"testing"
)

func TestGetSizeSuffixes(t *testing.T) {
	suffixes := GetSizeSuffixes()

	if len(suffixes) != 4 {
		t.Errorf("expected %d suffixes, but got %d", 4, len(suffixes))
	}

	for _, suffix := range suffixes {
		_, ok := SizeMap[suffix]
		if !ok {
			t.Errorf("expected %s in size map, but it is not a key.", suffix)
		}
	}
}

func TestCreate(t *testing.T) {
	var b bytes.Buffer
	bw := bufio.NewWriter(&b)

	expected := 25 * B

	err := Create(bw, int64(expected))
	if err != nil {
		t.Errorf("failed to create file %v", err)
	}

	actual := int64(len(b.Bytes()))
	if actual != int64(expected) {
		t.Errorf("expected length: %d but actual got: %d", 25, actual)
	}
}
