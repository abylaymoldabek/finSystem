package domain

import "time"

type Account struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

type Transaction struct {
	ID         int       `json:"id"`
	Value      float64   `json:"value"`
	AccountID  int       `json:"account_id"`
	Group      string    `json:"group"`
	Account2ID int       `json:"account2_id"`
	Date       time.Time `json:"date"`
}
