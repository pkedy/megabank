// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://api.evilcorp.com/terms/",
        "contact": {
            "name": "API Support",
            "url": "https://api.evilcorp.com/support",
            "email": "api@evilcorp.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts/v1/{id}": {
            "get": {
                "description": "Get the account balance by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Get the account balance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/accounts/v1/{id}/deposit": {
            "post": {
                "description": "Deposit into an account by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Deposit into an account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Amount request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AmountRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/accounts/v1/{id}/withdraw": {
            "post": {
                "description": "Withdraw from an account by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Withdraw from an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Amount request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.AmountRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/transfers/v1/{id}": {
            "post": {
                "description": "Transfer a specific amount of money from one account to another",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transfers"
                ],
                "summary": "Transfer a specific amount from one account to another",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Reference ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transfer request",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.TransferRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TransactionResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.AmountRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer",
                    "format": "int64"
                }
            }
        },
        "api.TransactionResult": {
            "type": "object",
            "properties": {
                "referenceId": {
                    "type": "string"
                },
                "transactionId": {
                    "type": "string"
                }
            }
        },
        "api.TransferRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "from": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8091",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MegaBank Account APIs",
	Description:      "This API was created using the Dapr SDK and is good but could be better.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
