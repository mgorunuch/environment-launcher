package config

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleShellCommand_Full() {
	cmd := ShellCommand{"/some/command", "-l", "hello"}
	fmt.Println(cmd.Full())
	// Output: /some/command -l hello
}

func TestShellCommand_Full(t *testing.T) {
	cases := []struct {
		name   string
		inp    ShellCommand
		output string
	}{
		{name: "no values", inp: ShellCommand{}, output: ""},
		{name: "base only", inp: ShellCommand{"some"}, output: "some"},
		{name: "base +1", inp: ShellCommand{"some", "-param"}, output: "some -param"},
		{name: "base +3", inp: ShellCommand{"a", "b", "c", "d"}, output: "a b c d"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := c.inp.Full()
			assert.Equal(t, c.output, res, "must equals")
		})
	}
}

func TestShellCommand_SplitBaseArgc(t *testing.T) {
	cases := []struct {
		name       string
		inp        ShellCommand
		outputBase string
		outputArgc []string
	}{
		{name: "empty", inp: ShellCommand{}, outputBase: "", outputArgc: nil},
		{name: "base only", inp: ShellCommand{"some"}, outputBase: "some", outputArgc: nil},
		{name: "base +1", inp: ShellCommand{"a", "b"}, outputBase: "a", outputArgc: []string{"b"}},
		{name: "base +3", inp: ShellCommand{"a", "b", "c", "d"}, outputBase: "a", outputArgc: []string{"b", "c", "d"}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resBase, resArgc := c.inp.SplitBaseArgc()
			assert.Equal(t, c.outputBase, resBase, "base must equals")
			assert.Equal(t, c.outputArgc, resArgc, "argc must equals")
		})
	}
}
