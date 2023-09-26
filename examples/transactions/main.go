package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
)

func main() {
	g := NewContainer()

	t1 := g.MustNewTeamStorage()
	t2 := g.MustNewTeamStorage()

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer func() {
		_ = tw.Flush()
	}()

	assertTransactionsSame(tw, "t1.Transaction == t1.UserStorage.Transaction", t1.Transaction, t1.UserStorage.Transaction)
	assertTransactionsSame(tw, "t1.Transaction == t1.ImageStorage.Transaction", t1.Transaction, t1.ImageStorage.Transaction)
	assertTransactionsNotSame(tw, "t1.Transaction != t2.Transaction", t1.Transaction, t2.Transaction, "each invocation `NewTeamStorage` initiates a new instance")
	assertTransactionsSame(tw, "t2.Transaction == t2.UserStorage.Transaction", t2.Transaction, t2.UserStorage.Transaction)
	assertTransactionsSame(tw, "t2.Transaction == t2.ImageStorage.Transaction", t2.Transaction, t2.ImageStorage.Transaction)
}

func assertTransactionsSame(w io.Writer, msg string, t1 Transaction, t2 Transaction) {
	_, _ = fmt.Fprint(w, msg+"\t")
	if t1 != t2 {
		panic("expected same transactions")
	}
	_, _ = fmt.Fprintln(w, "\t[✓]")
}

func assertTransactionsNotSame(w io.Writer, msg string, t1 Transaction, t2 Transaction, extra ...string) {
	_, _ = fmt.Fprint(w, msg+"\t")
	if t1 == t2 {
		panic("expected not same transactions")
	}
	_, _ = fmt.Fprintln(w, "\t[✓]\t"+strings.Join(extra, "\t"))
}
