package dummyfile

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
)

const (
	B        = 1
	KB int64 = 1 << (10 * iota)
	MB
	GB
	TB
)

var SizeMap = map[string]int64{
	"KB": KB,
	"MB": MB,
	"GB": GB,
	"TB": TB,
}

func GetSizeSuffixes() []string {
	suffixes := make([]string, 0, len(SizeMap))

	for key := range SizeMap {
		suffixes = append(suffixes, key)
	}

	return suffixes
}

func Create(file io.Writer, size int64) error {
	_, err := rand.Int(rand.Reader, big.NewInt(100))

	if err != nil {
		return fmt.Errorf("could not init a random value %v", err)
	}

	limitedReader := &io.LimitedReader{
		R: rand.Reader,
		N: size,
	}

	_, err = io.Copy(file, limitedReader)

	if err != nil {
		return fmt.Errorf("failed to copy data: %v", err)
	}

	return nil
}
