package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abylaymoldabek/finSystem/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initStudentRoutes(api *gin.RouterGroup) {
	api.POST("/accounts", h.createAccount)
	api.GET("/accounts", h.getAccounts)
	api.POST("/transactions", h.createTransaction)
	api.GET("/transactions/:account_id", h.getAccountTransactions)
}

// AddTags godoc
// @Summary Add a new account
// @Description Add a new account to the system
// @Tags tags
// @Accept json
// @Produce json
// @Param product body domain.Account true "Account object to be added"
// @Success 201 {string} string "Account added successfully"
// @Router /accounts [post]
func (h *Handler) createAccount(c *gin.Context) {
	var account domain.Account
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.services.Products.CreateAccount(c.Request.Context(), &account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account added successfully"})
}

// AddTags godoc
// @Summary Add a new transaction
// @Description Add a new transaction to the system
// @Tags tags
// @Accept json
// @Produce json
// @Param product body domain.Transaction true "Account object to be added"
// @Success 201 {string} string "Transaction added successfully"
// @Router /transactions [post]
func (h *Handler) createTransaction(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if transaction.Group == "transfer" && transaction.Account2ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "For transfer, account2_id must be provided"})
		return
	}

	err := h.services.Products.CreateTransaction(c.Request.Context(), &transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

// GetTags godoc
// @Summary Get all accounts
// @Description Get a list of all accounts in the system
// @Tags tags
// @Produce json
// @Success 200 {array} domain.Account "List of accounts"
// @Router /accounts [get]
func (h *Handler) getAccounts(c *gin.Context) {
	accounts, err := h.services.Products.GetAccounts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

// GetTags godoc
// @Summary Get all account transactions
// @Description Get a list of all account transactions in the system
// @Tags tags
// @Param account_id path int true "Account ID"
// @Produce json
// @Success 200 {array} domain.Transaction "List of account Transactions"
// @Router /transactions/{account_id} [get]
func (h *Handler) getAccountTransactions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("account_id"))
	fmt.Println(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}
	fmt.Println(id)

	accounts, err := h.services.Products.GetAccountTransactions(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}
