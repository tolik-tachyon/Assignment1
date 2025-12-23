package Wallet

import "fmt"

type Wallet struct {
	Balance      float64
	Transactions []float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
	w.Transactions = append(w.Transactions, amount)
}

func (w *Wallet) SpendMoney(amount float64) {
	if amount > w.Balance {
		fmt.Println("Not enough balance")
		return
	}
	w.Balance -= amount
	w.Transactions = append(w.Transactions, -amount)
}

func (w *Wallet) GetBalance() float64 {
	return w.Balance
}

func DemoWallet() {
	fmt.Println("\n--- WALLET DEMO ---")

	wallet := Wallet{}
	wallet.AddMoney(100)
	wallet.SpendMoney(30)

	fmt.Println("Balance:", wallet.GetBalance())
	fmt.Println("Transactions:", wallet.Transactions)
}
