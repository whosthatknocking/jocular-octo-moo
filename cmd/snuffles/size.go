package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	BYTE     = 1.0
	KILOBYTE = 1000 * BYTE
	MEGABYTE = 1000 * KILOBYTE
	GIGABYTE = 1000 * MEGABYTE
	TERABYTE = 1000 * GIGABYTE
	PETABYTE = 1000 * TERABYTE
)

type unitMap map[string]float64

var (
	decimalMap = unitMap{"K": KILOBYTE, "M": MEGABYTE, "G": GIGABYTE, "T": TERABYTE, "P": PETABYTE}
	sizeRegex  = regexp.MustCompile(`^([\d\.]+)\s{0,1}((?i)[kmgtp])(?i)b?$`)
)

// Return bytes from human-readable approximation of size
func FromHumanSize(size string) (int64, error) {
	return parseSizeToBytes(size, decimalMap)
}

// Parses the human-readable size string into the bytes amount it represents.
func parseSizeToBytes(sizeStr string, uMap unitMap) (int64, error) {
	matches := sizeRegex.FindStringSubmatch(sizeStr)
	if len(matches) != 3 {
		return -1, fmt.Errorf("invalid size: '%s'", sizeStr)
	}

	size, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return -1, err
	}

	unitPrefix := strings.ToUpper(matches[2])
	if mul, ok := uMap[unitPrefix]; ok {
		size *= mul
	}
	return int64(size), nil
}
