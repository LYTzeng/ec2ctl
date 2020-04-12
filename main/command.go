package main

import (
	"flag"
)

func ShowVersion() *Version {
	ver := &Version{
		fs: flag.NewFlagSet("version", flag.ContinueOnError),
	}
	return ver
}

func StopEC2() *Stop {
	stop := &Stop{
		fs: flag.NewFlagSet("stop", flag.ContinueOnError),
	}
	stop.fs.StringVar(&stop.instanceID, "i", "", "instance id")
	return stop
}

func StartEC2() *Start {
	start := &Start{
		fs: flag.NewFlagSet("start", flag.ContinueOnError),
	}
	start.fs.StringVar(&start.instanceID, "i", "", "instance id")
	return start
}

func ListEC2() *List {
	ls := &List{
		fs: flag.NewFlagSet("ls", flag.ContinueOnError),
	}
	ls.fs.StringVar(&ls.region, "r", "", "region")
	if ls.region == "" {
		ls.region = "all"
	}
	return ls
}
