package service

import (
	"context"
	"time"

	"github.com/abylaymoldabek/finSystem/internal/domain"
	"github.com/abylaymoldabek/finSystem/internal/repository"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewProductsService(repo repository.Transaction) *TransactionService {
	return &TransactionService{
		repo: repo,
	}
}

func (s *TransactionService) CreateAccount(ctx context.Context, account *domain.Account) error {
	return s.repo.CreateAccount(ctx, account)
}

func (s *TransactionService) GetAccounts(ctx context.Context) ([]*domain.Account, error) {
	return s.repo.GetAccounts(ctx)
}

func (s *TransactionService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) error {
	if transaction.Date.IsZero() {
		transaction.Date = time.Now()
	}
	return s.repo.CreateTransaction(ctx, transaction)
}

func (s *TransactionService) GetAccountTransactions(ctx context.Context, accountID int) ([]*domain.Transaction, error) {
	return s.repo.GetAccountTransactions(ctx, accountID)
}
