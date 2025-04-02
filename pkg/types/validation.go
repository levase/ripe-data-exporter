package types

import (
	"errors"
	"regexp"
	"strconv"
)

var (
	ErrASNEmpty      = errors.New("AS number cannot be empty")
	ErrASNInvalid    = errors.New("invalid AS number format")
	ErrASNOutOfRange = errors.New("AS number out of valid range")
)

func ValidateASN(asn string) error {
	if asn == "" {
		return ErrASNEmpty
	}

	re := regexp.MustCompile(`^(AS)?(\d+)$`)
	matches := re.FindStringSubmatch(asn)
	if matches == nil {
		return ErrASNInvalid
	}

	numStr := matches[2]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return ErrASNInvalid
	}

	if num < 1 || num > 4294967295 {
		return ErrASNOutOfRange
	}

	return nil
}
