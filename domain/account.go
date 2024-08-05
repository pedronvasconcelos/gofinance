package domain

import (
	"fmt"
	"math/rand"

	"github.com/shopspring/decimal"
)

type Account struct {
	Id        int
	FirstName string
	LastName  string
	Number    string
	Balance   decimal.Decimal
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:        rand.Int(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    fmt.Sprintf("%d", rand.Intn(1000000000)),
	}
}
