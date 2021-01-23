package config

import (
	"bytes"
)

type ShellCommand []string

func (sc ShellCommand) Full() string {
	if len(sc) == 0 {
		return ""
	}

	var s bytes.Buffer
	s.WriteString(sc[0])

	for i := 1; i < len(sc); i++ {
		s.WriteRune(' ')
		s.WriteString(sc[i])
	}

	return s.String()
}

func (sc ShellCommand) SplitBaseArgc() (string, []string) {
	if len(sc) == 0 {
		return "", nil
	}

	baseCmd := sc[0]
	var arguments []string
	if len(sc) > 1 {
		arguments = sc[1:]
	}

	return baseCmd, arguments
}
