package cmd_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func assertBuff(t *testing.T, r io.Reader, expected string) {
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
		return
	}

	assert.Equal(t, expected, string(out))
}

func assertCmd(t *testing.T, cmd *cobra.Command, args []string, out string, err bool) {
	buff := bytes.NewBufferString("")
	buffErr := bytes.NewBufferString("")
	cmd.SetOut(buff)
	cmd.SetErr(buffErr)
	cmd.SetArgs(args)

	cmdErr := cmd.Execute()
	if err {
		assert.Error(t, cmdErr)
	} else {
		assert.NoError(t, cmdErr)
	}

	assertBuff(t, buff, out)
}

type cmdScenario struct {
	cmd   *cobra.Command
	args  string
	out   string
	error bool
}

func runCmdScenarios(t *testing.T, scenarios ...cmdScenario) {
	for i, tmp := range scenarios {
		s := tmp
		var args []string
		if s.args != "" {
			args = strings.Split(s.args, " ")
		}
		t.Run(fmt.Sprintf("Scenario #%d", i), func(t *testing.T) {
			assertCmd(
				t,
				s.cmd,
				args,
				s.out,
				s.error,
			)
		})
	}
}
