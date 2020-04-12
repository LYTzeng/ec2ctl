package main

// Runner decides which subcommand to run
type Runner interface {
	Init([]string) error
	Run() string
	Name() string
}
