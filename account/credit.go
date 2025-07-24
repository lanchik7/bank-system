package account

import (
	"errors"
)

type CreditAccount struct{
	id int
	balance float64
	creditLimit float64
}

var lastID int
func generateID() int {
    lastID++
    return lastID
}

func NewCreditAccount(initialBalance float64, creditLimit float64) (*CreditAccount, error) {
	if initialBalance < -creditLimit {
        return nil, errors.New("начальный баланс превышает кредитный лимит")
    }
	id := generateID()

	return &CreditAccount{
		id: id,
		balance: initialBalance,
		creditLimit: creditLimit,
	}, nil
}

func (c *CreditAccount) Deposit(amount float64) error {
	if (amount <= 0){
		return errors.New("сумма должна быть положительной")
	}
	c.balance += amount
	return nil
}

func (c *CreditAccount) Withdraw(amount float64) error{
	if (amount<=0) {
		return errors.New("сумма должна быть положительной")
	}
	if(c.balance - amount < -c.creditLimit){
		return errors.New("сумма должна быть положительной")
	}
	c.balance -= amount
	return nil
}