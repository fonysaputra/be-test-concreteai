package handlers

import (
	"encoding/json"
	"fmt"
	"ms-go-gin/config"
	"ms-go-gin/internal/domain"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Account struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewAccount(cfg *config.Config, db *gorm.DB) *Account {
	return &Account{cfg, db}
}

// @Summary Get List My Payment Account
// @Description Get list my payment account with token
// @Tags account
// @Accept  json
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} domain.SampleRespSuccessGeneral "successful get list my account"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/account [get]
func (app *Account) List(c *gin.Context) {
	var (
		db                  = app.db
		mListPaymentAccount []domain.PaymentAccount
		dataUser            domain.UserResponse
	)

	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	fmt.Println("userData :", cast.ToString(userData))
	json.Unmarshal([]byte(cast.ToString(userData)), &dataUser)
	db.Session(&gorm.Session{PrepareStmt: false}).Where("user_id = ?", dataUser.ID).Find(&mListPaymentAccount)

	c.JSON(http.StatusOK, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusOK), Message: "success", Data: mListPaymentAccount})
}

// @Summary Create Payment Account
// @Description Create payment account and every create payment account balance is set default to 100$
// @Tags account
// @Accept  json
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param   paymentAccountRequest  body    domain.PaymentAccountRequest  true  "Payment Account Request Body"
// @Success 201 {object} domain.PaymentAccountRequest "successful create payment account"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/account/create [post]
func (app *Account) Create(c *gin.Context) {
	var (
		request             domain.PaymentAccountRequest
		db                  = app.db
		mListPaymentAccount []domain.PaymentAccount
		dataUser            domain.UserResponse
		dataCurrency        domain.Currencies
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

	resultCurr := db.Session(&gorm.Session{PrepareStmt: false}).Where("currency_code = ?", strings.ToUpper(request.Currancy)).Find(&dataCurrency)
	if resultCurr.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusNotFound), Message: "not_found"})
		return
	}

	resultListAccount := db.Session(&gorm.Session{PrepareStmt: false}).Where("account_number = ?", request.AccountNumber).Find(&mListPaymentAccount)

	if resultListAccount.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "account_already_registered"})
		return
	}

	savePayment := &domain.PaymentAccount{UserID: cast.ToString(dataUser.ID), AccountType: request.AccountType, CurrencyId: dataCurrency.ID, Balance: 100, AccountNumber: request.AccountNumber}
	db.Create(&savePayment)
	c.JSON(http.StatusCreated, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusCreated), Message: "created", Data: savePayment})
}
