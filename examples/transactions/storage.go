package main

import (
	"math/rand"
)

type Transaction interface {
}

type UserStorage struct {
	Transaction Transaction
}

type ImageStorage struct {
	Transaction Transaction
}

type TeamStorage struct {
	Transaction  Transaction
	UserStorage  *UserStorage
	ImageStorage *ImageStorage
}

type MockTransaction struct {
	ID int
}

func NewMockTransaction() *MockTransaction {
	return &MockTransaction{ID: rand.Int()}
}
