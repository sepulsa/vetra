package viper

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"testing"
)

func TestRun(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "github.com/spf13/viper", "with_alias_import", "with_dot_import", "with_viper")
}
