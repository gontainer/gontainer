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

package runner

import (
	"io"
	"strings"
)

const (
	rowWidth  = 60
	checkMark = "[✓]"
	xMark     = "[⨉]"
)

type printer interface {
	Println(string)
	PrintAlignedLn(left string, extra ...string)
}

type indenter interface {
	Indent(string)
	EndIndent()
}

type Printer struct {
	writer  io.Writer
	indents []string
}

func NewPrinter(w io.Writer) *Printer {
	return &Printer{
		writer: w,
	}
}

func (p *Printer) Indent(s string) {
	p.indents = append(p.indents, s)
}

func (p *Printer) EndIndent() {
	p.indents = p.indents[:len(p.indents)-1]
}

func (p *Printer) Println(s string) {
	_, err := p.writer.Write([]byte(strings.Join(p.indents, "") + s + "\n"))
	if err != nil {
		panic(err.Error())
	}
}

func (p *Printer) PrintAlignedLn(left string, extra ...string) {
	right := ""
	if len(extra) > 0 {
		right = extra[0]
		extra = extra[1:]
	}
	m := strings.Repeat(
		"·",
		rowWidth-len([]rune(left+right+strings.Join(p.indents, ""))),
	)
	line := append([]string{left, m, right}, extra...)
	p.Println(strings.Join(line, ""))
}
