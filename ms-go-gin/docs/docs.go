// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/account": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get list my payment account with token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Get List My Payment Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful get list my account",
                        "schema": {
                            "$ref": "#/definitions/domain.SampleRespSuccessGeneral"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/account/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create payment account and every create payment account balance is set default to 100$",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "Create Payment Account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Payment Account Request Body",
                        "name": "paymentAccountRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.PaymentAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "successful create payment account",
                        "schema": {
                            "$ref": "#/definitions/domain.PaymentAccountRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/login": {
            "post": {
                "description": "Log in an existing user with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Log in a user",
                "parameters": [
                    {
                        "description": "Login Request Body",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/domain.SampleRespSuccessGeneral"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/profile": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user profile by validating JWT token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile data",
                        "schema": {
                            "$ref": "#/definitions/domain.SampleRespSuccessGeneral"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/signup": {
            "post": {
                "description": "Register a new user with an email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign up a new user",
                "parameters": [
                    {
                        "description": "Register Request Body",
                        "name": "registerRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Signup successful",
                        "schema": {
                            "$ref": "#/definitions/domain.SampleRespSuccessGeneral"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/transactions/send": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Transactions Send money from account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Send",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Transactions Send Request Body",
                        "name": "sendRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.TrxRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "transactions success",
                        "schema": {
                            "$ref": "#/definitions/domain.TrxRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/transactions/withdrawal": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Withdraw money from account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Withdrawal",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Withdrawal Send Request Body",
                        "name": "trxRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.TrxRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "transactions success",
                        "schema": {
                            "$ref": "#/definitions/domain.TrxRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/transactions/{account_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all transactions for an account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get all transactions for an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "get list transactions",
                        "schema": {
                            "$ref": "#/definitions/domain.TrxRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "domain.PaymentAccountRequest": {
            "type": "object",
            "required": [
                "accountNumber"
            ],
            "properties": {
                "accountNumber": {
                    "type": "string",
                    "example": "2777625252525252"
                },
                "accountType": {
                    "type": "string",
                    "example": "debit"
                },
                "currency": {
                    "type": "string",
                    "example": "USD"
                }
            }
        },
        "domain.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "user@example.com"
                },
                "password": {
                    "type": "string",
                    "example": "password123"
                }
            }
        },
        "domain.SampleRespSuccessGeneral": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "200"
                },
                "message": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "domain.TrxRequest": {
            "type": "object",
            "required": [
                "accountNumber",
                "to_address"
            ],
            "properties": {
                "accountNumber": {
                    "type": "string",
                    "example": "2777625252525252"
                },
                "amount": {
                    "type": "string",
                    "example": "5"
                },
                "to_address": {
                    "type": "string",
                    "example": "2777625252525252"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
