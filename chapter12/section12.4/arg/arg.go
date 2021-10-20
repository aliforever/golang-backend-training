package arg

import (
	"flag"
	"os"
)

var (
	Values []string
)

var (
	Query string
)

var (
	Verbose                     bool
	VeryVerbose                 bool
	VeryVeryVerbose             bool
	VeryVeryVeryVerbose         bool
	VeryVeryVeryVeryVerbose     bool
	VeryVeryVeryVeryVeryVerbose bool
	DeploymentEnvironment       string
)

var (
	Help bool
)

func init() {
	flag.StringVar(&Query, "q", "", "search query")

	flag.BoolVar(&Verbose, "v", false, "verbose logs outputted")
	flag.BoolVar(&VeryVerbose, "vv", false, "very verbose logs outputted")
	flag.BoolVar(&VeryVeryVerbose, "vvv", false, "very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVerbose, "vvvv", false, "very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVerbose, "vvvvv", false, "very very very very verbose logs outputted")
	flag.BoolVar(&VeryVeryVeryVeryVeryVerbose, "vvvvvv", false, "very very very very very verbose logs outputted")

	flag.StringVar(&DeploymentEnvironment, "help", "DEV", "Deployment Environment (DEV, QA, STAGING, PROD)")

	flag.BoolVar(&Help, "help", false, "outputs help message")

	flag.Parse()

	Values = flag.Args()

	// --help
	if Help {
		flag.PrintDefaults()
		os.Exit(0)
	}
}
