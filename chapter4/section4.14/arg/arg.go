package arg

import (
	"flag"
)

var LogLevel uint8 = 0

func init() {
	var (
		flagLevel1 bool
		flagLevel2 bool
		flagLevel3 bool
		flagLevel4 bool
		flagLevel5 bool
		flagLevel6 bool
	)

	flag.BoolVar(&flagLevel1, "v", false, "-v")
	flag.BoolVar(&flagLevel2, "vv", false, "-vv")
	flag.BoolVar(&flagLevel3, "vvv", false, "-vvv")
	flag.BoolVar(&flagLevel4, "vvvv", false, "-vvvv")
	flag.BoolVar(&flagLevel5, "vvvvv", false, "-vvvvv")
	flag.BoolVar(&flagLevel6, "vvvvvv", false, "-vvvvvv")
	flag.Parse()

	if flagLevel1 {
		LogLevel = 1
	} else if flagLevel2 {
		LogLevel = 2
	} else if flagLevel3 {
		LogLevel = 3
	} else if flagLevel4 {
		LogLevel = 4
	} else if flagLevel5 {
		LogLevel = 5
	} else if flagLevel6 {
		LogLevel = 6
	}
}
