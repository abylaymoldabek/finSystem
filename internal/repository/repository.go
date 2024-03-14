package repository

import (
	"context"
	"database/sql"

	"github.com/abylaymoldabek/finSystem/internal/domain"
)

type Transaction interface {
	CreateAccount(ctx context.Context, account *domain.Account) error
	GetAccounts(ctx context.Context) ([]*domain.Account, error)
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) error
	GetAccountTransactions(ctx context.Context, accountID int) ([]*domain.Transaction, error)
}

type Repository struct {
	Products Transaction
}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{
		Products: NewProductsRepo(db),
	}
}
