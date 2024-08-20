package handlers

import (
	"encoding/json"
	"fmt"
	"ms-go-gin/config"
	"ms-go-gin/internal/domain"
	"ms-go-gin/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type Auth struct {
	cfg *config.Config
}

func NewAuth(cfg *config.Config) *Auth {
	return &Auth{cfg}
}

// @Summary Sign up a new user
// @Description Register a new user with an email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   registerRequest  body    domain.RegisterRequest  true  "Register Request Body"
// @Success 200 {object} domain.SampleRespSuccessGeneral "Signup successful"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/signup [post]
func (app *Auth) SignUp(c *gin.Context) {

	var (
		userCredentials domain.RegisterRequest
		data            domain.ResDataSuccess
	)

	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "bad_request"})
		return
	}
	fmt.Println("Supabase URL:", app.cfg.SupabaseUrl)
	fmt.Println("Supabase API Key:", app.cfg.SupabaseApikey)

	resp, err := helper.GetRestyClient(app.cfg).R().EnableTrace().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(userCredentials).
		Post("/auth/v1/signup")

	fmt.Println("Response Status:", resp.Status())
	fmt.Println("Response Body:", resp.String())

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(resp.StatusCode(), &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "bad_request"})
		return
	}

	json.Unmarshal(resp.Body(), &data)

	c.JSON(http.StatusOK, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusOK), Message: "success", Data: data})

}

// @Summary Log in a user
// @Description Log in an existing user with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   loginRequest  body    domain.LoginRequest  true  "Login Request Body"
// @Success 200 {object} domain.SampleRespSuccessGeneral "Login successful"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Router /api/v1/login [post]
func (app *Auth) Login(c *gin.Context) {
	var (
		userCredentials domain.LoginRequest
		data            domain.ResDataSuccess
	)

	if err := c.ShouldBindJSON(&userCredentials); err != nil {
		fmt.Println("error", err.Error())
		c.JSON(http.StatusBadRequest, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "bad_request"})
		return
	}

	resp, err := helper.GetRestyClient(app.cfg).R().
		EnableTrace().
		SetHeaders(map[string]string{"Content-Type": "application/json"}).
		SetBody(userCredentials).
		Post("/auth/v1/token?grant_type=password")

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(resp.StatusCode(), &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusBadRequest), Message: "bad_request"})
		return
	}

	json.Unmarshal(resp.Body(), &data)

	c.JSON(http.StatusOK, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusOK), Message: "success", Data: data})
}

// profile handler returns user profile data
// @Summary Get user profile
// @Description Get user profile by validating JWT token
// @Tags Profile
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} domain.SampleRespSuccessGeneral "Profile data"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /api/v1/profile [get]
func (app *Auth) Profile(c *gin.Context) {
	// Retrieve the user data stored in the context by the middleware
	userData, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, &domain.HttpResponseGeneral{Code: cast.ToString(http.StatusInternalServerError), Message: "server_error"})
		return
	}

	// Return the user profile
	c.JSON(http.StatusOK, gin.H{"profile": userData})
}
