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

package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	goversion "github.com/caarlos0/go-version"
	"github.com/gontainer/gontainer/internal/cmd"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

var (
	version    = ""
	commit     = ""
	isGitDirty = ""
	date       = ""
	builtBy    = ""

	//go:embed logo.txt
	logo string
)

const (
	name        = "gontainer"
	description = "Build DI containers based on YAML configuration"
	url         = "https://github.com/gontainer/gontainer"
)

func main() {
	bv := buildVersion()

	rootCmd := cobra.Command{
		Use:     name,
		Short:   description,
		Version: bv.GitVersion,
	}

	rootCmd.SetVersionTemplate(bv.String())

	rootCmd.AddCommand(
		cmd.NewBuildCmd(
			bv.GitVersion,
			fmt.Sprintf("%s %s-%s %s", bv.GitVersion, bv.GitCommit, bv.GitTreeState, bv.BuildDate),
		),
	)

	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildVersion() goversion.Info {
	return goversion.GetVersionInfo(
		goversion.WithAppDetails(
			name,
			description,
			url,
		),
		goversion.WithASCIIName(logo),
		func(i *goversion.Info) {
			if commit != "" {
				i.GitCommit = commit
			}
			switch isGitDirty {
			case "true":
				i.GitTreeState = "dirty"
			case "false":
				i.GitTreeState = "clean"
			}
			if date != "" {
				i.BuildDate = date
			}
			if version != "" {
				i.GitVersion = version
			}
			if strings.HasPrefix(i.GitVersion, "v") && semver.IsValid(i.GitVersion) {
				i.GitVersion = strings.TrimPrefix(i.GitVersion, "v")
			}
			if builtBy != "" {
				i.BuiltBy = builtBy
			}
		},
	)
}
