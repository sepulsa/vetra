package main

import (
	"github.com/sepulsa/vetra/analysis/passes/viper"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(viper.Analyzer)
}
