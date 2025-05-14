package histogram

import (
	"fmt"
	"sort"
	"strings"
)

// Histogram with simple string bins
type Histogram struct {
	Bins map[string]int
}

func New() *Histogram {
	return &Histogram{
		Bins: make(map[string]int),
	}
}

// Add a bin (or increase the count by 1)
func (h *Histogram) Add(bin string) {
	h.AddInc(bin, 1)
}

// AddInc a bin (or increase the count by increase)
func (h *Histogram) AddInc(bin string, increase int) {
	h.Bins[bin] += increase
}

// Get the count for a bin
func (h *Histogram) Get(bin string) int {
	return h.Bins[bin]
}

// Print the histogram to stdout
func (h *Histogram) Print() {
	keys := make([]string, 0, len(h.Bins))
	for k := range h.Bins {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, bin := range keys {
		count := h.Bins[bin]
		label := fmt.Sprintf("[%s]", bin)
		bar := strings.Repeat("=", count)
		fmt.Printf("%-3s %3d %s\n", label, count, bar)
	}
}
