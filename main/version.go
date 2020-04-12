package main

import "flag"

type Version struct {
	fs *flag.FlagSet
}

func (ver *Version) Name() string {
	return ver.fs.Name()
}

func (cmd *Version) Init(args []string) error {
	return cmd.fs.Parse(args)
}

func (cmd *Version) Run() string {
	return "Version 0.0"
}
