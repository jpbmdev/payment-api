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
