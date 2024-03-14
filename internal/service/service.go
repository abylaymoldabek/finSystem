package service

import (
	"context"

	"github.com/abylaymoldabek/finSystem/internal/domain"
	"github.com/abylaymoldabek/finSystem/internal/repository"
)

type Products interface {
	CreateAccount(ctx context.Context, account *domain.Account) error
	GetAccounts(ctx context.Context) ([]*domain.Account, error)
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) error
	GetAccountTransactions(ctx context.Context, accountID int) ([]*domain.Transaction, error)
}

type Service struct {
	Products Products
}

type Deps struct {
	Repos *repository.Repository
}

func NewServices(deps Deps) *Service {
	productsService := NewProductsService(deps.Repos.Products)

	return &Service{
		Products: productsService,
	}
}
