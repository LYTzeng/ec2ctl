package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"
	"text/tabwriter"
)

type List struct {
	fs     *flag.FlagSet
	region string
}

func (ls *List) Name() string {
	return ls.fs.Name()
}

func (ls *List) Init(args []string) error {
	return ls.fs.Parse(args)
}

func (ls *List) Run() string {
	ec2Collector := newCollector()
	var lsInstance []ec2Instance
	if ls.region == "all" {
		lsInstance = ec2Collector.collectAll()
	} else if found := find(allRegions, ls.region); found {
		lsInstance = ec2Collector.collectInRegion(ls.region)
	} else {
		return "No such region!"
	}
	var b bytes.Buffer
	w := new(tabwriter.Writer)
	w.Init(&b, 0, 0, 1, ' ', 0)
	fmt.Fprintln(w, "INSTANCE NAME\t ID\t TYPE\t STATE\t REGION")
	for _, instance := range lsInstance {
		fmt.Fprintln(w, strings.Join([]string{instance.Name, instance.InstanceID, instance.InstanceType, instance.State, instance.Region}, "\t "))
	}
	w.Flush()
	return b.String()
}

func find(slice [18]string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
