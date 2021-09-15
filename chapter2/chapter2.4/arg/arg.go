package arg

import "flag"

func ProcessFlags(name *string, shhh *bool) {
	flag.StringVar(name, "name", "Joe", "--name=Joe")
	flag.BoolVar(shhh, "shhh", false, "--shhh")
	flag.Parse()
}
