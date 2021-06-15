package operator

import (
	"regexp/syntax"
	"testing"

	"github.com/waflab/waflab/autogen/utils"
)

func TestReggen(t *testing.T) {
	tests := []struct {
		expression string
		output     string
	}{
		{`(abc\b|a)\sf`, `abc f`},
		{`(abc |a)\ba`, `abc a`},
	}

	for _, tc := range tests {
		t.Run(tc.expression, func(t *testing.T) {
			utils.SetRandomSeed(41)
			re, err := syntax.Parse(tc.expression, syntax.PerlX)
			if err != nil {
				t.Fatal(err)
			}
			out := string(generate(re))
			if out != tc.output {
				t.Fatalf("expect: %s, got: %s\n", tc.output, out)
			}
		})
	}
}
