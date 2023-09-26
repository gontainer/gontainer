package pkg

import (
	"math/rand"
	"strconv"
)

type Transaction struct {
	ID string
}

func NewTransaction() *Transaction {
	return &Transaction{
		ID: strconv.FormatInt(rand.Int63(), 16),
	}
}

type Accounts struct {
	Transaction *Transaction
}

type Users struct {
	Accounts    *Accounts
	Transaction *Transaction
}
