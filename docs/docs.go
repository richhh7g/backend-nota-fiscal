// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Richhh7g",
            "url": "https://github.com/richhh7g",
            "email": "richhh7g@protonmail.com"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/richhh7g/backend-nota-fiscal/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/nota-fiscal": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1",
                    "Invoices"
                ],
                "summary": "Create invoice",
                "parameters": [
                    {
                        "description": "Invoice",
                        "name": "invoice",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateInvoiceBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/CreateInvoiceResponse"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/nota-fiscal/{chave}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1",
                    "Invoices"
                ],
                "summary": "Get invoice",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Chave",
                        "name": "chave",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetInvoiceResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateInvoiceBody": {
            "type": "object",
            "required": [
                "chave",
                "cnpj",
                "data_emissao",
                "data_recebimento"
            ],
            "properties": {
                "chave": {
                    "type": "string",
                    "example": "1234567890123456789012345678901234567890123456789012345678901234"
                },
                "cnpj": {
                    "type": "string",
                    "example": "12345678901234"
                },
                "data_emissao": {
                    "type": "string",
                    "example": "2022-08-01T10:00:00Z"
                },
                "data_recebimento": {
                    "type": "string",
                    "example": "2022-08-01T10:00:00Z"
                }
            }
        },
        "CreateInvoiceResponse": {
            "type": "object",
            "properties": {
                "chave": {
                    "type": "string",
                    "example": "1234567890123456789012345678901234567890123456789012345678901234"
                },
                "cnpj": {
                    "type": "string",
                    "example": "12345678901234"
                },
                "data_emissao": {
                    "type": "string",
                    "example": "2022-08-01T10:00:00Z"
                },
                "data_recebimento": {
                    "type": "string",
                    "example": "2022-08-01T10:00:00Z"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "12345678-9abc-def0-1234-56789abcdef0"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "errors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "GetInvoiceResponse": {
            "type": "object",
            "properties": {
                "chave": {
                    "type": "string",
                    "example": "12345678901234567890123456789012345678901234"
                },
                "cnpj": {
                    "type": "string",
                    "example": "12345678901234"
                },
                "data_emissao": {
                    "type": "string",
                    "example": "2022-08-01T10:00:00Z"
                },
                "data_recebimento": {
                    "type": "string",
                    "example": "2022-08-01T10:00:00Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "NFE API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
