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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Jaithun Mahira",
            "email": "swiggyb1035@datascience.manipal.edu"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "This request will return 200 OK if server is up..",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "To check if the service is running or not.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        },
        "/transaction/{customerId}": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Get the total transaction amount credited to the customer based on customer ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction Service"
                ],
                "summary": "Get transaction amount for a customer.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer Id",
                        "name": "customerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Transaction"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will add transaction amount to a customer by customer ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transacion Service"
                ],
                "summary": "Adds transaction amount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer Id",
                        "name": "customerId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Transaction details",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Transaction"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "number"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "number"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Transaction": {
            "type": "object",
            "required": [
                "amount",
                "customer_id"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "customer_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3005",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Transaction Service API",
	Description:      "Admin can add and view transaction amount for a customer based on custmer ID",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
