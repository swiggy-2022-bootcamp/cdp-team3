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
        "contact": {
            "name": "Chahat Bhatia",
            "email": "chahatbhatia2014@gmail.com"
        },
        "license": {
            "name": "Apache 2.0"
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
                    ""
                ],
                "summary": "To check if the service is running or not.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/shippingaddress": {
            "post": {
                "description": "This Handler allow user to create new Shipping Address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shipping Address"
                ],
                "summary": "Create Shipping Address",
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
        "/shippingaddress/:id": {
            "get": {
                "description": "This Handle returns shippingAddress given id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shipping Address"
                ],
                "summary": "Get Shipping Address by id",
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
            },
            "put": {
                "description": "This Handle Update shippingAddress given id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shipping Address"
                ],
                "summary": "Update Shipping Address",
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
            },
            "delete": {
                "description": "This Handle deletes Delete Shipping Address given sid",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shipping Address"
                ],
                "summary": "Delete Shipping Address",
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
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3002",
	BasePath:         "/categories-service/api",
	Schemes:          []string{},
	Title:            "CDP TEAM 3 - Categories module",
	Description:      "This microservice is for categories service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
