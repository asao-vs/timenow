package main

import (
	"github.com/asao-vs/timenow"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(timenow.Analyzer) }
