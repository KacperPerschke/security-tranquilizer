package img

import (
	"cmp"
	"errors"
	"slices"
)

// CDS is a container for excerpt from https://en.wikipedia.org/wiki/Computer_display_standard
type CDS struct {
	Width  int
	Height int
}

func (ss CDS) Capacity() int {
	return ss.Width * ss.Height
}

func (ss CDS) WillAccommodate(cap int) bool {
	return ss.Capacity() > cap
}

var prefSizes = []CDS{
	CDS{
		Width:  160,
		Height: 120,
	},
	CDS{
		Width:  160,
		Height: 128,
	},
	CDS{
		Width:  160,
		Height: 144,
	},
	CDS{
		Width:  240,
		Height: 160,
	},
	CDS{
		Width:  320,
		Height: 240,
	},
	CDS{
		Width:  640,
		Height: 200,
	},
	CDS{
		Width:  480,
		Height: 272,
	},
	CDS{
		Width:  640,
		Height: 256,
	},
	CDS{
		Width:  512,
		Height: 342,
	},
	CDS{
		Width:  640,
		Height: 350,
	},
	CDS{
		Width:  720,
		Height: 348,
	},
	CDS{
		Width:  720,
		Height: 350,
	},
	CDS{
		Width:  640,
		Height: 400,
	},
	CDS{
		Width:  640,
		Height: 480,
	},
	CDS{
		Width:  720,
		Height: 480,
	},
	CDS{
		Width:  1024,
		Height: 768,
	},
	CDS{
		Width:  1280,
		Height: 720,
	},
	CDS{
		Width:  1152,
		Height: 864,
	},
	CDS{
		Width:  1366,
		Height: 768,
	},
	CDS{
		Width:  1280,
		Height: 960,
	},
	CDS{
		Width:  1440,
		Height: 900,
	},
	CDS{
		Width:  1600,
		Height: 900,
	},
	CDS{
		Width:  1400,
		Height: 1050,
	},
	CDS{
		Width:  1680,
		Height: 1050,
	},
	CDS{
		Width:  1600,
		Height: 1200,
	},
	CDS{
		Width:  1920,
		Height: 1080,
	},
	CDS{
		Width:  2048,
		Height: 1080,
	},
	CDS{
		Width:  1920,
		Height: 1200,
	},
	CDS{
		Width:  2048,
		Height: 1152,
	},
	CDS{
		Width:  1920,
		Height: 1280,
	},
	CDS{
		Width:  1920,
		Height: 1440,
	},
	CDS{
		Width:  2160,
		Height: 1440,
	},
	CDS{
		Width:  2048,
		Height: 1536,
	},
	CDS{
		Width:  2560,
		Height: 1440,
	},
	CDS{
		Width:  2560,
		Height: 1600,
	},
	CDS{
		Width:  2880,
		Height: 1440,
	},
	CDS{
		Width:  2960,
		Height: 1440,
	},
	CDS{
		Width:  3440,
		Height: 1440,
	},
	CDS{
		Width:  2736,
		Height: 1824,
	},
	CDS{
		Width:  2880,
		Height: 1800,
	},
	CDS{
		Width:  2560,
		Height: 2048,
	},
	CDS{
		Width:  3024,
		Height: 1964,
	},
	CDS{
		Width:  3000,
		Height: 2000,
	},
	CDS{
		Width:  3840,
		Height: 1600,
	},
	CDS{
		Width:  3200,
		Height: 2048,
	},
	CDS{
		Width:  3200,
		Height: 2400,
	},
	CDS{
		Width:  3456,
		Height: 2234,
	},
	CDS{
		Width:  3840,
		Height: 2160,
	},
	CDS{
		Width:  4096,
		Height: 2160,
	},
	CDS{
		Width:  5120,
		Height: 2160,
	},
	CDS{
		Width:  4096,
		Height: 3072,
	},
	CDS{
		Width:  4500,
		Height: 3000,
	},
	CDS{
		Width:  5120,
		Height: 2880,
	},
	CDS{
		Width:  5120,
		Height: 3200,
	},
	CDS{
		Width:  5120,
		Height: 4096,
	},
	CDS{
		Width:  7680,
		Height: 3200,
	},
	CDS{
		Width:  6400,
		Height: 4096,
	},
	CDS{
		Width:  6400,
		Height: 4800,
	},
	CDS{
		Width:  7680,
		Height: 4320,
	},
	CDS{
		Width:  7680,
		Height: 4800,
	},
	CDS{
		Width:  10240,
		Height: 4320,
	},
}

// resBySize returns the smallest element in the prefSizes providing Capacity
// that accomodate expected number of bytes.
func resBySize(expected int) (CDS, error) {
	areTwoCDSItemsInOrder := func(a, b CDS) int {
		return cmp.Compare(a.Capacity(), b.Capacity())
	}
	if !slices.IsSortedFunc(prefSizes, areTwoCDSItemsInOrder) {
		slices.SortFunc(prefSizes, areTwoCDSItemsInOrder)
	}
	for _, e := range prefSizes {
		if e.WillAccommodate(expected) {
			return e, nil
		}
	}
	return CDS{}, errors.New("haven't found meeting preffered size")
}
