// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/2fa/confirm-google-totp": {
            "get": {
                "description": "Google-Authenticator Confirmation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "2FA"
                ],
                "summary": "Google-Authenticator Confirmation route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "X-Request-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "uint",
                        "description": " ",
                        "name": "otp",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/2fa/confirm-sms-otp": {
            "get": {
                "description": "SMS Confirmation",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "2FA"
                ],
                "summary": "SMS Confirmation route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "X-Request-Id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "uint",
                        "description": " ",
                        "name": "otp",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/2fa/send-sms-otp": {
            "get": {
                "description": "SMS sender",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "2FA"
                ],
                "summary": "SMS sender route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "X-Request-Id",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.OTP"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Authentication",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Authentication route",
                "parameters": [
                    {
                        "format": "json",
                        "description": "authentication",
                        "name": "authentication",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/2fa-init": {
            "get": {
                "description": "TwoFAInit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "TwoFAInit route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/conversion": {
            "post": {
                "description": "Conversion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "Conversion route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "format": "json",
                        "description": "conversion",
                        "name": "conversion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ConversionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ConversionResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/get-accounts": {
            "get": {
                "description": "GetAccounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "GetAccounts route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAccountsResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/get-all-amount": {
            "get": {
                "description": "GetAllAmount",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "GetAllAmount route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "json",
                        "description": " ",
                        "name": "currency",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetAllAmountsResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/get-limit-conversion": {
            "get": {
                "description": "GetLimitConversion",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "GetLimitConversion route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "json",
                        "description": " ",
                        "name": "currency",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetLimitConversion"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/get-rates": {
            "get": {
                "description": "GetExchangeRates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "GetExchangeRates route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetExchangeRatesResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/merchant/get-transactions": {
            "get": {
                "description": "GetTransactions",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Merchant"
                ],
                "summary": "GetTransactions route",
                "parameters": [
                    {
                        "type": "string",
                        "format": "UUID",
                        "description": " ",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "json",
                        "description": " ",
                        "name": "dateFrom",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "json",
                        "description": " ",
                        "name": "dateTo",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetTransactionResponse"
                        }
                    },
                    "303": {
                        "description": "See Other",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "merchant.Accounts": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "balance": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "typeAcc": {
                    "type": "string"
                }
            }
        },
        "merchant.Rates": {
            "type": "object",
            "properties": {
                "buyRate": {
                    "type": "string"
                },
                "currency": {
                    "type": "string"
                },
                "sellRate": {
                    "type": "string"
                }
            }
        },
        "merchant.Transactions": {
            "type": "object",
            "properties": {
                "accDtNumber": {
                    "type": "string"
                },
                "accKtNumber": {
                    "type": "string"
                },
                "dateProcess": {
                    "type": "string"
                },
                "docId": {
                    "type": "string"
                },
                "docState": {
                    "type": "string"
                },
                "nazn": {
                    "type": "string"
                },
                "recipientName": {
                    "type": "string"
                },
                "senderName": {
                    "type": "string"
                },
                "summa": {
                    "type": "string"
                },
                "transType": {
                    "type": "string"
                }
            }
        },
        "models.ConversionRequest": {
            "type": "object",
            "properties": {
                "accFrom": {
                    "type": "string"
                },
                "accTo": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "dest": {
                    "type": "string"
                }
            }
        },
        "models.ConversionResponse": {
            "type": "object",
            "properties": {
                "docID": {
                    "type": "string"
                },
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                }
            }
        },
        "models.ErrorModel": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorDescription": {
                    "type": "string"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorDescription": {
                    "type": "string"
                },
                "response": {}
            }
        },
        "models.GetAccountsResponse": {
            "type": "object",
            "properties": {
                "accounts": {
                    "$ref": "#/definitions/merchant.Accounts"
                },
                "merchName": {
                    "type": "string"
                },
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                }
            }
        },
        "models.GetAllAmountsResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "merchName": {
                    "type": "string"
                },
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                }
            }
        },
        "models.GetExchangeRatesResponse": {
            "type": "object",
            "properties": {
                "base": {
                    "type": "string"
                },
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                },
                "rates": {
                    "$ref": "#/definitions/merchant.Rates"
                }
            }
        },
        "models.GetLimitConversion": {
            "type": "object",
            "properties": {
                "currentLimit": {
                    "type": "number"
                },
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                }
            }
        },
        "models.GetTransactionResponse": {
            "type": "object",
            "properties": {
                "merchName": {
                    "type": "string"
                },
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                },
                "trnList": {
                    "$ref": "#/definitions/merchant.Transactions"
                }
            }
        },
        "models.Login": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "maxLength": 9,
                    "minLength": 9
                }
            }
        },
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "$ref": "#/definitions/models.ErrorModel"
                },
                "response": {
                    "$ref": "#/definitions/models.twoFa"
                }
            }
        },
        "models.OTP": {
            "type": "object",
            "properties": {
                "expirition": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "totalSent": {
                    "type": "integer"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "refreshToken": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.googleAuth": {
            "type": "object",
            "properties": {
                "isActive": {
                    "type": "boolean"
                }
            }
        },
        "models.smsOTP": {
            "type": "object",
            "properties": {
                "isActive": {
                    "type": "boolean"
                }
            }
        },
        "models.twoFa": {
            "type": "object",
            "properties": {
                "googleAuthenticator": {
                    "$ref": "#/definitions/models.googleAuth"
                },
                "smsOtp": {
                    "$ref": "#/definitions/models.smsOTP"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8070",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "B2B Example API",
	Description:      "This is a B2B-API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
