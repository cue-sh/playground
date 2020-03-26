package main

import (
	"fmt"
	"testing"
)

var testTable = []struct {
	In     input
	Fn     function
	Out    output
	InVal  string
	OutVal string
	Err    string
}{
	{inputCUE, functionEval, outputCUE, "", "\n", ""},
	{inputCUE, functionEval, outputCUE, "a: b: 5\na: c: 4", "a: {\n\tb: 5\n\tc: 4\n}\n", ""},
	{inputCUE, functionEval, outputJSON, "test: 5", "{\n    \"test\": 5\n}\n", ""},
	{inputJSON, functionEval, outputCUE, "{\n \"test\": 5\n}\n", "test: 5", ""},
}

func TestHandleCUECompile(t *testing.T) {
	for _, tv := range testTable {
		desc := fmt.Sprintf("handleCUECompile(%q, %q, %q, %q)", tv.In, tv.Fn, tv.Out, tv.InVal)
		out, err := handleCUECompile(tv.In, tv.Fn, tv.Out, tv.InVal)
		if tv.Err != "" {
			if err != nil {
				if err.Error() != tv.Err {
					t.Fatalf("%v: expected error string %q; got %q", desc, tv.Err, err)
				}
			} else {
				t.Fatalf("%v: expected error, did not see one. Output was %q", desc, out)
			}
		} else {
			if err != nil {
				t.Fatalf("%v: got unexpected error: %v", desc, err)
			} else if out != tv.OutVal {
				t.Fatalf("%v: expected output %q: got %q", desc, tv.OutVal, out)
			}
		}
	}
}
