// TODO: answer here
package account

type Account struct {
	Name    string
	Balance int
}

func (x Account) GetBalance() int {
	// TODO: answer here
	return x.Balance
}

func (a *Account) Deposit(amount int) {
	// TODO: answer here
	a.Balance += amount
}
