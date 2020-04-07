package main

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"cuelang.org/go/cmd/cue/cmd"
	"cuelang.org/go/cue/load"
)

type function string

const (
	functionEval         function = "eval"
	functionEvalConcrete function = "eval-c"
)

type input string

const (
	inputCUE  input = "cue"
	inputJSON input = "json"
	inputYaml input = "yaml"
)

type output string

const (
	outputCUE  output = output(inputCUE)
	outputJSON output = output(inputJSON)
	outputYaml output = output(inputYaml)
)

func handleCUECompile(in input, fn function, out output, inputVal string) (string, error) {
	// TODO implement more functions
	switch fn {
	case functionEval, functionEvalConcrete:
	default:
		return "", fmt.Errorf("function %q is not implemented", fn)
	}

	args := []string{"eval"}

	switch out {
	case outputCUE, outputJSON, outputYaml:
		args = append(args, "--out", string(out))
	default:
		return "", fmt.Errorf("unknown ouput type: %v", out)
	}

	switch in {
	case inputCUE, inputJSON, inputYaml:
		args = append(args, string(in), "/input.cue")
	default:
		return "", fmt.Errorf("unknown input type: %v", in)
	}

	c, err := cmd.New(args)
	root := "/"
	c.Dir = &root
	c.ModuleRoot = &root
	c.Overlay = map[string]load.Source{
		"/cue.mod/module.cue": load.FromString(`module: "cue.playground"`),
		"/input.cue":          load.FromString(inputVal),
	}
	if err != nil {
		return "", fmt.Errorf("failed to parse command from args [%v]: %v", strings.Join(args, " "), err)
	}
	var outBuf bytes.Buffer
	c.SetOutput(&outBuf)
	if err := c.Run(context.Background()); err != nil {
		// Return this error "naked" because it contains the actual
		// error we want to show to the user
		return "", err
	}
	return outBuf.String(), nil
}
