package main

import "errors"

// cds is a container for excerpt from https://en.wikipedia.org/wiki/Computer_display_standard
type cds struct {
	Width    int
	Height   int
	Capacity int
}

var prefSizes = []cds{
	cds{
		Width:    160,
		Height:   120,
		Capacity: 19200,
	},
	cds{
		Width:    160,
		Height:   128,
		Capacity: 20480,
	},
	cds{
		Width:    160,
		Height:   144,
		Capacity: 23040,
	},
	cds{
		Width:    240,
		Height:   160,
		Capacity: 38400,
	},
	cds{
		Width:    320,
		Height:   240,
		Capacity: 76800,
	},
	cds{
		Width:    640,
		Height:   200,
		Capacity: 128000,
	},
	cds{
		Width:    480,
		Height:   272,
		Capacity: 130560,
	},
	cds{
		Width:    640,
		Height:   256,
		Capacity: 163840,
	},
	cds{
		Width:    512,
		Height:   342,
		Capacity: 175104,
	},
	cds{
		Width:    640,
		Height:   350,
		Capacity: 224000,
	},
	cds{
		Width:    720,
		Height:   348,
		Capacity: 250560,
	},
	cds{
		Width:    720,
		Height:   350,
		Capacity: 252000,
	},
	cds{
		Width:    640,
		Height:   400,
		Capacity: 256000,
	},
	cds{
		Width:    640,
		Height:   480,
		Capacity: 307200,
	},
	cds{
		Width:    720,
		Height:   480,
		Capacity: 345600,
	},
	cds{
		Width:    1024,
		Height:   768,
		Capacity: 786432,
	},
	cds{
		Width:    1280,
		Height:   720,
		Capacity: 921600,
	},
	cds{
		Width:    1152,
		Height:   864,
		Capacity: 995328,
	},
	cds{
		Width:    1366,
		Height:   768,
		Capacity: 1049088,
	},
	cds{
		Width:    1280,
		Height:   960,
		Capacity: 1228800,
	},
	cds{
		Width:    1440,
		Height:   900,
		Capacity: 1296000,
	},
	cds{
		Width:    1600,
		Height:   900,
		Capacity: 1440000,
	},
	cds{
		Width:    1400,
		Height:   1050,
		Capacity: 1470000,
	},
	cds{
		Width:    1680,
		Height:   1050,
		Capacity: 1764000,
	},
	cds{
		Width:    1600,
		Height:   1200,
		Capacity: 1920000,
	},
	cds{
		Width:    1920,
		Height:   1080,
		Capacity: 2073600,
	},
	cds{
		Width:    2048,
		Height:   1080,
		Capacity: 2211840,
	},
	cds{
		Width:    1920,
		Height:   1200,
		Capacity: 2304000,
	},
	cds{
		Width:    2048,
		Height:   1152,
		Capacity: 2359296,
	},
	cds{
		Width:    1920,
		Height:   1280,
		Capacity: 2457600,
	},
	cds{
		Width:    1920,
		Height:   1440,
		Capacity: 2764800,
	},
	cds{
		Width:    2160,
		Height:   1440,
		Capacity: 3110400,
	},
	cds{
		Width:    2048,
		Height:   1536,
		Capacity: 3145728,
	},
	cds{
		Width:    2560,
		Height:   1440,
		Capacity: 3686400,
	},
	cds{
		Width:    2560,
		Height:   1600,
		Capacity: 4096000,
	},
	cds{
		Width:    2880,
		Height:   1440,
		Capacity: 4147200,
	},
	cds{
		Width:    2960,
		Height:   1440,
		Capacity: 4262400,
	},
	cds{
		Width:    3440,
		Height:   1440,
		Capacity: 4953600,
	},
	cds{
		Width:    2736,
		Height:   1824,
		Capacity: 4990464,
	},
	cds{
		Width:    2880,
		Height:   1800,
		Capacity: 5184000,
	},
	cds{
		Width:    2560,
		Height:   2048,
		Capacity: 5242880,
	},
	cds{
		Width:    3024,
		Height:   1964,
		Capacity: 5939136,
	},
	cds{
		Width:    3000,
		Height:   2000,
		Capacity: 6000000,
	},
	cds{
		Width:    3840,
		Height:   1600,
		Capacity: 6144000,
	},
	cds{
		Width:    3200,
		Height:   2048,
		Capacity: 6553600,
	},
	cds{
		Width:    3200,
		Height:   2400,
		Capacity: 7680000,
	},
	cds{
		Width:    3456,
		Height:   2234,
		Capacity: 7720704,
	},
	cds{
		Width:    3840,
		Height:   2160,
		Capacity: 8294400,
	},
	cds{
		Width:    4096,
		Height:   2160,
		Capacity: 8847360,
	},
	cds{
		Width:    5120,
		Height:   2160,
		Capacity: 11059200,
	},
	cds{
		Width:    4096,
		Height:   3072,
		Capacity: 12582912,
	},
	cds{
		Width:    4500,
		Height:   3000,
		Capacity: 13500000,
	},
	cds{
		Width:    5120,
		Height:   2880,
		Capacity: 14745600,
	},
	cds{
		Width:    5120,
		Height:   3200,
		Capacity: 16384000,
	},
	cds{
		Width:    5120,
		Height:   4096,
		Capacity: 20971520,
	},
	cds{
		Width:    7680,
		Height:   3200,
		Capacity: 24576000,
	},
	cds{
		Width:    6400,
		Height:   4096,
		Capacity: 26214400,
	},
	cds{
		Width:    6400,
		Height:   4800,
		Capacity: 30720000,
	},
	cds{
		Width:    7680,
		Height:   4320,
		Capacity: 33177600,
	},
	cds{
		Width:    7680,
		Height:   4800,
		Capacity: 36864000,
	},
	cds{
		Width:    10240,
		Height:   4320,
		Capacity: 44236800,
	},
}

// resBySize the function finds the smallest element in the prefSizes
// providing `Capacity` greater than `expected`.
// The search method assumes that the `prefSizes` is sorted
// by the contents of the `Capacity` key.
func resBySize(expected int) (cds, error) {
	for _, e := range prefSizes {
		if e.Capacity >= expected {
			return e, nil
		}
	}
	return cds{}, errors.New("haven't found meeting preffered size")
}
