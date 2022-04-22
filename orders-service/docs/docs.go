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
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will fetch all the orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Fetch all the orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
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
        "/orders/invoice/{orderId}": {
            "post": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will generate an invoice for order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Generate invoice for a particular Order by Order ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
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
        "/orders/status/{status}": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Get all the orders in the application based on the order status for admin to view.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Get orders based on order status",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Status",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
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
        "/orders/user/{userId}": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Get order details of a customer based on Customer ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Get orders of a customer based on customer ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "userId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
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
        "/orders/{orderId}": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "Get order details based on Order ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Get order based on order ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
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
            "put": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will update the order status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Update Order Status by Order ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order Status",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.OrderStatus"
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
            },
            "delete": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will delete a particular order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Delete Order by Order ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
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
        "/orders/{orderId}/order_status": {
            "get": {
                "security": [
                    {
                        "Bearer Token": []
                    }
                ],
                "description": "This request will fetch details of order status",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders Service"
                ],
                "summary": "Get Order Status by Order ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order Id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.OrderStatus"
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
        "models.Order": {
            "type": "object",
            "required": [
                "dateTime"
            ],
            "properties": {
                "customerId": {
                    "type": "string"
                },
                "dateTime": {
                    "type": "string"
                },
                "orderId": {
                    "type": "string"
                },
                "orderedProducts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.OrderedProduct"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "models.OrderStatus": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "models.OrderedProduct": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
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
	Host:             "localhost:3004",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Orders Service API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
