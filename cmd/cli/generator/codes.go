package generator

import (
	"physical-pin-code-generator/cmd/cli/histogram"
	"physical-pin-code-generator/cmd/cli/storage"
)

func GenerateMissingCodes(storage *storage.Storage, h *histogram.Histogram) {
	for i, code := range storage.Codes {
		if code.Code == "" {
			GeneratePin(&storage.Codes[i], h)
		}
	}
}
