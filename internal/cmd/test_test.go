// Copyright (c) 2023 Bartłomiej Krukowski
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is furnished
// to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
