package repository

import (
	"context"
	"database/sql"

	"github.com/abylaymoldabek/finSystem/internal/domain"
)

type ProductsRepo struct {
	db *sql.DB
}

func NewProductsRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{
		db: db,
	}
}

func (r *ProductsRepo) CreateAccount(ctx context.Context, account *domain.Account) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO accounts (name, balance) VALUES ($1, $2)", account.Name, account.Balance)
	return err
}

func (r *ProductsRepo) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, balance FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var accounts []*domain.Account
	for rows.Next() {
		account := &domain.Account{}
		err := rows.Scan(&account.ID, &account.Name, &account.Balance)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r *ProductsRepo) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO transactions (value, account_id, group_type, account2_id) VALUES ($1, $2, $3, $4)", transaction.Value, transaction.AccountID, transaction.Group, transaction.Account2ID)
	return err
}

func (r *ProductsRepo) GetAccountTransactions(ctx context.Context, accountID int) ([]*domain.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, value, account_id, group_type, account2_id, date FROM transactions WHERE account_id = $1", accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transactions []*domain.Transaction
	for rows.Next() {
		transaction := &domain.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.Value, &transaction.AccountID, &transaction.Group, &transaction.Account2ID, &transaction.Date)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
