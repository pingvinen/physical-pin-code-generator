package generator

import (
	"math/rand"
	"physical-pin-code-generator/cmd/cli/histogram"
	"physical-pin-code-generator/cmd/cli/storage"
)

func GeneratePin(code *storage.CodesEntry, h *histogram.Histogram) {
	newCode := ""

	for i := 0; i < code.Length; i++ {
		// Get the allowed keys for the next key in the code.
		// This will be the least used keys according to the histogram.
		allowedKeys := getAllowedKeys(h)

		// Pick a random key from the allowed keys.
		nextKey := allowedKeys[rand.Intn(len(allowedKeys))]

		// Update the histogram.
		h.AddInc(nextKey, code.Frequency)

		// Update the code.
		newCode += nextKey
	}

	code.Code = newCode
}

func getAllowedKeys(h *histogram.Histogram) []string {
	allowedKeys := make([]string, 0)

	for bin, count := range h.Bins {
		if count == h.Lowest {
			allowedKeys = append(allowedKeys, bin)
		}
	}

	return allowedKeys
}
