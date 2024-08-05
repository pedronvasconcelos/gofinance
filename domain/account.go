package domain

import (
	"fmt"
	"math/rand"

	"github.com/shopspring/decimal"
)

type Account struct {
	Id        int             `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Number    string          `json:"number"`
	Balance   decimal.Decimal `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		Id:        rand.Int(),
		FirstName: firstName,
		LastName:  lastName,
		Number:    fmt.Sprintf("%d", rand.Intn(1000000000)),
	}
}
