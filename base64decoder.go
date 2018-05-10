package config

import (
	"encoding/base64"

	"github.com/pkg/errors"
)

// DecodeBase64Values will decode list of base64 encoded values to list of decoded []byte
func DecodeBase64Values(values []string, results []*[]byte) error {
	if len(values) != len(results) {
		return errors.New("number of values and results must be equal")
	}

	for i, value := range values {
		result, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return errors.Wrapf(err, "failed to decode %d value", i)
		}
		*(results[i]) = result
	}
	return nil
}
