// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/decoration/delete": {
            "delete": {
                "consumes": [
                    "application/text"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cake Decorations"
                ],
                "summary": "Delete a given cake decoration by its article",
                "parameters": [
                    {
                        "description": "Article of a cake decoration to be deleted",
                        "name": "Article",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/decoration/edit": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cake Decorations"
                ],
                "summary": "Edit a given decoration",
                "parameters": [
                    {
                        "description": "scheme of cake decoration",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CakeDecoration"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.CakeDecoration"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/decoration/get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cake Decorations"
                ],
                "summary": "Get a list of all cake decorations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.CakeDecoration"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ingredient/delete": {
            "delete": {
                "consumes": [
                    "application/text"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Delete a given ingredient by its marking",
                "parameters": [
                    {
                        "description": "Article of an ingredient to be deleted",
                        "name": "Article",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ingredient/edit": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Edit a given ingredient",
                "parameters": [
                    {
                        "description": "scheme of ingredient",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Ingredient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Ingredient"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/ingredient/get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ingredients"
                ],
                "summary": "Get a list of all ingredients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Ingredient"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorisation"
                ],
                "summary": "Login of user",
                "parameters": [
                    {
                        "description": "scheme of login",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/order/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "description": "scheme of order request",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.OrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Tokens"
                ],
                "summary": "Refresh access token",
                "parameters": [
                    {
                        "description": "scheme of login",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authorisation"
                ],
                "summary": "Signup of clients",
                "parameters": [
                    {
                        "description": "scheme of login",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignupRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignupResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tooling/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Toolings"
                ],
                "summary": "Create a new tooling",
                "parameters": [
                    {
                        "description": "scheme of login",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ToolingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Tooling"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tooling/delete": {
            "delete": {
                "consumes": [
                    "application/text"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Toolings"
                ],
                "summary": "Delete a tooling",
                "parameters": [
                    {
                        "description": "ID of a tooling to be deleted",
                        "name": "Tooling_ID",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tooling/edit": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Toolings"
                ],
                "summary": "Update a tooling",
                "parameters": [
                    {
                        "description": "scheme of login",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ToolingRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Tooling"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/tooling/get": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Toolings"
                ],
                "summary": "Get a list of all toolings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.ToolingResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CakeDecoration": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "article": {
                    "type": "string"
                },
                "costPrice": {
                    "type": "number"
                },
                "image": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "supplier": {
                    "$ref": "#/definitions/domain.Supplier"
                },
                "supplierName": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "domain.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "domain.Ingredient": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "article": {
                    "type": "string"
                },
                "costPrice": {
                    "type": "number"
                },
                "gost": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "ingredientType": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "packing": {
                    "type": "string"
                },
                "specs": {
                    "type": "string"
                },
                "supplier": {
                    "$ref": "#/definitions/domain.Supplier"
                },
                "supplierName": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "domain.Item": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                }
            }
        },
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "domain.Order": {
            "type": "object",
            "properties": {
                "assignedManager": {
                    "$ref": "#/definitions/domain.User"
                },
                "assignedManagerName": {
                    "type": "string"
                },
                "date": {
                    "type": "string"
                },
                "examples": {
                    "type": "string"
                },
                "expectedFulfilmentDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "item": {
                    "$ref": "#/definitions/domain.Item"
                },
                "itemName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orderer": {
                    "$ref": "#/definitions/domain.User"
                },
                "ordererName": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "domain.OrderRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "examples": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "manager": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "orderer": {
                    "type": "string"
                },
                "size": {
                    "type": "string"
                }
            }
        },
        "domain.RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "domain.SignupRequest": {
            "type": "object",
            "required": [
                "login",
                "password"
            ],
            "properties": {
                "fullName": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.SignupResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "domain.Supplier": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "deliveryTime": {
                    "description": "в часах",
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.Tooling": {
            "type": "object",
            "properties": {
                "marking": {
                    "type": "string"
                },
                "properties": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/domain.ToolingType"
                },
                "typeName": {
                    "type": "string"
                }
            }
        },
        "domain.ToolingRequest": {
            "type": "object",
            "properties": {
                "acquireTime": {
                    "type": "string"
                },
                "amount": {
                    "type": "integer"
                },
                "decayLevel": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "supplier": {
                    "$ref": "#/definitions/domain.Supplier"
                },
                "type": {
                    "$ref": "#/definitions/domain.ToolingType"
                }
            }
        },
        "domain.ToolingResponse": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "amount": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "domain.ToolingType": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "fullName": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "photoURL": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "Описание работы сервера",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
