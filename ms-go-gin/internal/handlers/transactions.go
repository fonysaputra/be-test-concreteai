package handlers

import (
	"encoding/json"
	"fmt"
	"ms-go-gin/config"
	"ms-go-gin/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Transactions struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewTransactions(cfg *config.Config, db *gorm.DB) *Transactions {
	return &Transactions{cfg, db}
}

// @Summary Send
// @Description Transactions Send money from account
// @Tags transactions
// @Accept  json
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param   sendRequest  body    domain.TrxRequest  true  "Transactions Send Request Body"
// @Success 200 {object} domain.TrxRequest "transactions success"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/transactions/send [post]
func (app *Transactions) Send(c *gin.Context) {
	var (
		db           = app.db
		dataUser     domain.UserResponse
		request      domain.TrxRequest
		mTransaction domain.Transaction
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "bad_request", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	fmt.Println("userData :", cast.ToString(userData))
	json.Unmarshal([]byte(cast.ToString(userData)), &dataUser)

	// Validate and process the transaction
	if err := db.Transaction(func(tx *gorm.DB) error {
		var account domain.PaymentAccount
		if err := tx.Where("account_number = ? and user_id = ?", request.AccountNumber, dataUser.ID).First(&account).Error; err != nil {
			//c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
			return err
		}

		// Update balance
		if cast.ToFloat64(account.Balance) < cast.ToFloat64(request.Amount) {
			c.JSON(http.StatusBadRequest, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "insufficient_balance"})

		}
		account.Balance -= cast.ToFloat64(request.Amount)
		if err := tx.Save(&account).Error; err != nil {
			return err
		}
		mTransaction = domain.Transaction{AccountId: account.ID, Amount: cast.ToFloat64(request.Amount), ToAddress: request.ToAddress, Status: "pending", Description: "success transfer to " + request.ToAddress}
		// Create the transaction record
		if err := tx.Create(&mTransaction).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error", Data: map[string]interface{}{"error": err.Error()}})

		return
	}
	c.JSON(http.StatusCreated, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusCreated), Message: "success", Data: mTransaction})
}

// @Summary Withdrawal
// @Description Withdraw money from account
// @Tags transactions
// @Accept  json
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param   trxRequest  body    domain.TrxRequest  true  "Withdrawal Send Request Body"
// @Success 200 {object} domain.TrxRequest "transactions success"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/transactions/withdrawal [post]
func (app *Transactions) Withdrawal(c *gin.Context) {
	var (
		db           = app.db
		dataUser     domain.UserResponse
		request      domain.TrxRequest
		mTransaction domain.Transaction
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "bad_request", Data: map[string]interface{}{"error": err.Error()}})
		return
	}

	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	fmt.Println("userData :", cast.ToString(userData))
	json.Unmarshal([]byte(cast.ToString(userData)), &dataUser)

	// Validate and process the transaction
	if err := db.Transaction(func(tx *gorm.DB) error {
		var account domain.PaymentAccount
		if err := tx.Where("account_number = ? and user_id = ?", request.AccountNumber, dataUser.ID).First(&account).Error; err != nil {
			return err
		}

		// Update balance

		account.Balance += cast.ToFloat64(request.Amount)
		if err := tx.Save(&account).Error; err != nil {
			return err
		}
		mTransaction = domain.Transaction{AccountId: account.ID, Amount: cast.ToFloat64(request.Amount), ToAddress: request.ToAddress, Status: "pending", Description: "success withdrawal from " + request.ToAddress}
		// Create the transaction record
		if err := tx.Create(&mTransaction).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error", Data: map[string]interface{}{"error": err.Error()}})

		return
	}
	c.JSON(http.StatusCreated, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusCreated), Message: "success", Data: mTransaction})
}

// @Summary Get all transactions for an account
// @Description Get all transactions for an account
// @Tags transactions
// @Accept  json
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param account_id path string true "Account ID"
// @Success 200 {object} domain.TrxRequest "get list transactions"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/transactions/{account_id} [get]
func (app *Transactions) GetTransactions(c *gin.Context) {
	var (
		db           = app.db
		dataUser     domain.UserResponse
		transactions []domain.Transaction
		account      domain.PaymentAccount
	)

	accountID := c.Param("account_id")

	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	fmt.Println("userData :", cast.ToString(userData))
	json.Unmarshal([]byte(cast.ToString(userData)), &dataUser)

	if err := db.Where("account_number = ? and user_id = ?", accountID, dataUser.ID).First(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	if err := db.Where("account_id = ? ", account.ID).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	c.JSON(http.StatusOK, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusOK), Message: "success", Data: transactions})
}
