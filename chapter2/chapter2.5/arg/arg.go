package arg

import "flag"

var (
	Name string
	Shhh bool
)

func init() {
	flag.StringVar(&Name, "name", "Joe", "--name=Joe")
	flag.BoolVar(&Shhh, "shhh", false, "--shhh")
	flag.Parse()
}
