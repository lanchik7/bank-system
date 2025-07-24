package clients

import "bank-system/account"

type Client struct {
	ID       int
	FullName string
	Accounts []account.Account
}

func (c *Client) AddAccount(acc account.Account) {
	c.Accounts = append(c.Accounts, acc)
}

func (c *Client) GetBalance() float64 {
	totalBalance := 0.0
	for _, acc := range c.Accounts{
		totalBalance += acc.GetBalance()
	}

	return totalBalance
}
