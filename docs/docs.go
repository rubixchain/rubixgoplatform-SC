// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "murali.c@ensurity.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/commit-data-token": {
            "post": {
                "description": "This API will create data token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Create Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Batch ID",
                        "name": "batchID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/create-data-token": {
            "post": {
                "description": "This API will create data token",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Create Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User/Entity ID",
                        "name": "UserID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "User/Entity Info",
                        "name": "UserInfo",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Committer DID",
                        "name": "CommitterDID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "Batch ID",
                        "name": "BacthID",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "File Info is json string {",
                        "name": "FileInfo",
                        "in": "formData"
                    },
                    {
                        "type": "file",
                        "description": "File to be committed",
                        "name": "FileContent",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/fetch-smart-contract": {
            "post": {
                "description": "This API will deploy smart contract",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Deploy Smart Contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "smartContractToken",
                        "name": "smartContractToken",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/generate-smart-contract": {
            "post": {
                "description": "This API will deploy smart contract",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Smart Contract"
                ],
                "summary": "Deploy Smart Contract",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "location of binary code hash",
                        "name": "binaryCodePath",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "location of raw code hash",
                        "name": "rawCodePath",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "location of Schema code hash",
                        "name": "schemaFilePath",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-account-info": {
            "get": {
                "description": "For a mentioned DID, check the account balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Check account balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-comment": {
            "get": {
                "description": "Retrieves the details of a transaction based on its comment.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by Transcation Comment",
                "operationId": "get-by-comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment to identify the transaction",
                        "name": "Comment",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-did": {
            "get": {
                "description": "Retrieves the details of a transaction based on dID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by dID",
                "operationId": "get-by-did",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID of sender/receiver",
                        "name": "DID",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Filter by role as sender or receiver",
                        "name": "Role",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-by-txnId": {
            "get": {
                "description": "Retrieves the details of a transaction based on its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get transaction details by Transcation ID",
                "operationId": "get-txn-details-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The ID of the transaction to retrieve",
                        "name": "txnID",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/get-data-token": {
            "get": {
                "description": "This API will get all data token belong to the did",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Data Tokens"
                ],
                "summary": "Get Data Token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "DID",
                        "name": "did",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/initiate-rbt-transfer": {
            "post": {
                "description": "Initiates a transfer of RBT tokens from one account to another.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Initiate RBT transfer",
                "operationId": "initiate-rbt-transfer",
                "parameters": [
                    {
                        "description": "Transfer input parameters",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.RBTTransferRequestSwaggoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/signature-response": {
            "post": {
                "description": "This API is used to supply the password for the node along with the ID generated when Initiate RBT transfer is called.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Signature Response",
                "operationId": "signature-response",
                "parameters": [
                    {
                        "description": "Transfer input parameters",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/server.inputData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        },
        "/api/start": {
            "get": {
                "description": "It will setup the core if not done before",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Basic"
                ],
                "summary": "Start Core",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.BasicResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.BasicResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "boolean"
                }
            }
        },
        "server.RBTTransferRequestSwaggoInput": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "receiver": {
                    "type": "string"
                },
                "sender": {
                    "type": "string"
                },
                "tokenCOunt": {
                    "type": "number"
                },
                "type": {
                    "type": "integer"
                }
            }
        },
        "server.inputData": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "mode": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "SessionToken": {
            "type": "apiKey",
            "name": "Session-Token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.9",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Rubix Core",
	Description:      "Rubix core API to control & manage the node.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
