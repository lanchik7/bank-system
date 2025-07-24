package account

import (
	"errors"
	"math/rand"
	"time"
)

type CreditAccount struct{
	id int
	balance float64
	creditLimit float64
}

func NewCreditAccount(initialBalance float64, creditLimit float64) (*CreditAccount, error) {
	rand.Seed(time.Now().UnixNano())
	if initialBalance < -creditLimit {
        return nil, errors.New("начальный баланс превышает кредитный лимит")
    }
	id := rand.Intn(1000000) + 1

	return &CreditAccount{
		id: id,
		balance: initialBalance,
		creditLimit: creditLimit,
	}, nil
}