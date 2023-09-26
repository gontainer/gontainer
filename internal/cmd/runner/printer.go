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
