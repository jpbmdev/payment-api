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
        "/loan": {
            "get": {
                "description": "Get loans, you can pass a start date, and a end date and the endpoint will find all loans STARTED in that range, if no params are passed this will return all loans, this endpoint supports a very simple pagination where you can select the page and the pageSize",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loan"
                ],
                "summary": "Get Loans",
                "parameters": [
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "to",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "int valid",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Loan"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            },
            "post": {
                "description": "Create loan, Full detail on the readme",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loan"
                ],
                "summary": "Create Loan",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "CreateLoanDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateLoanDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Loan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        },
        "/loan/{id}": {
            "get": {
                "description": "Get loan by Id",
                "tags": [
                    "loan"
                ],
                "summary": "Get Loan by Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Loan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        },
        "/loan/{id}/debt": {
            "get": {
                "description": "Get Debt of a single Loan, you can pass a date to check the debt on that specific time, if no date is passed this endpoint will return the entire debt",
                "tags": [
                    "loan"
                ],
                "summary": "Get Debt of a single Loan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "string valid",
                        "name": "date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoanDebt"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        },
        "/loan/{id}/payment": {
            "get": {
                "description": "Get Payments by loan Id",
                "tags": [
                    "loan"
                ],
                "summary": "Get Payments by loan Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
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
                                "$ref": "#/definitions/models.Payment"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            },
            "put": {
                "description": "Add Payment To Loan",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loan"
                ],
                "summary": "Add Payment To Loan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "payload",
                        "name": "AddPaymentToLoanDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddPaymentToLoanDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Payment"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        },
        "/target-schema": {
            "get": {
                "description": "Get lists of Target Schemas",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "target-schema"
                ],
                "summary": "Get Target Schemas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TargetSchemaSwagger"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        },
        "/target-schema/test-tree": {
            "post": {
                "description": "Test Target Schema Decision Tree output with Cant and AmountTotal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "target-schema"
                ],
                "summary": "Test Target Schema Decision Tree",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "DecisionTreeInputs",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.DecisionTreeInputs"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TargetParams"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get lists of users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a user with his name",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "CreateUserDto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateUserDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.SucessfullOperation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.FailedOperation"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddPaymentToLoanDto": {
            "type": "object",
            "required": [
                "amount",
                "date"
            ],
            "properties": {
                "amount": {
                    "type": "number"
                },
                "date": {
                    "type": "string"
                }
            }
        },
        "models.CreateLoanDto": {
            "type": "object",
            "required": [
                "amount",
                "startDate",
                "term",
                "userId"
            ],
            "properties": {
                "amount": {
                    "type": "number",
                    "minimum": 1
                },
                "startDate": {
                    "type": "string"
                },
                "term": {
                    "type": "integer",
                    "minimum": 1
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.CreateUserDto": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "models.DecisionTreeInputs": {
            "type": "object",
            "properties": {
                "amountTotal": {
                    "type": "number"
                },
                "cant": {
                    "type": "number"
                }
            }
        },
        "models.FailedOperation": {
            "type": "object",
            "properties": {
                "internalCode": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Loan": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "debt": {
                    "type": "number"
                },
                "endDate": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "loanHistory": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.LoanHistory"
                    }
                },
                "quota": {
                    "type": "number"
                },
                "rate": {
                    "type": "number"
                },
                "startDate": {
                    "type": "string"
                },
                "targetName": {
                    "type": "string"
                },
                "targetSchemaId": {
                    "type": "string"
                },
                "term": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "models.LoanDebt": {
            "type": "object",
            "properties": {
                "debt": {
                    "type": "number"
                }
            }
        },
        "models.LoanHistory": {
            "type": "object",
            "properties": {
                "accumulated": {
                    "type": "number"
                },
                "monthDebt": {
                    "type": "number"
                },
                "monthEnd": {
                    "type": "string"
                },
                "monthStart": {
                    "type": "string"
                },
                "paymentId": {
                    "type": "string"
                }
            }
        },
        "models.Payment": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "loanId": {
                    "type": "string"
                }
            }
        },
        "models.SucessfullOperation": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.TargetParams": {
            "type": "object",
            "properties": {
                "max": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rate": {
                    "type": "number"
                }
            }
        },
        "models.TargetSchemaSwagger": {
            "type": "object",
            "properties": {
                "desicionTree": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tree"
                    }
                },
                "id": {
                    "type": "string"
                },
                "targets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TargetParams"
                    }
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.Tree": {
            "type": "object",
            "properties": {
                "content": {},
                "headers": {
                    "type": "object",
                    "additionalProperties": true
                },
                "id": {
                    "type": "integer"
                },
                "key": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "operator": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "parentId": {
                    "type": "integer"
                },
                "value": {}
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Payment Simple API",
	Description:      "This is a sample server to manage payments",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
