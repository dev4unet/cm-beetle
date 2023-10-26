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
            "name": "API Support",
            "url": "http://cloud-barista.github.io",
            "email": "contact-to-cloud-barista@googlegroups.com"
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
        "/config": {
            "get": {
                "description": "List all configs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System environment"
                ],
                "summary": "List all configs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.RestGetAllConfigResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "post": {
                "description": "Create or Update config (SPIDER_REST_URL, DRAGONFLY_REST_URL, ...)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System environment"
                ],
                "summary": "Create or Update config",
                "parameters": [
                    {
                        "description": "Key and Value for configuration",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.ConfigReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ConfigInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "delete": {
                "description": "Init all configs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System environment"
                ],
                "summary": "Init all configs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            }
        },
        "/config/{configId}": {
            "get": {
                "description": "Get config",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System environment"
                ],
                "summary": "Get config",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Config ID",
                        "name": "configId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ConfigInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "delete": {
                "description": "Init config",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System environment"
                ],
                "summary": "Init config",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Config ID",
                        "name": "configId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.ConfigInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Check Beetle is alive",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System management"
                ],
                "summary": "Check Beetle is alive",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            }
        },
        "/httpVersion": {
            "get": {
                "description": "Checks and logs the HTTP version of the incoming request to the server console.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Admin] System management"
                ],
                "summary": "Check HTTP version of incoming request",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            }
        },
        "/ns": {
            "get": {
                "description": "List all namespaces or namespaces' ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Namespace] Namespace management"
                ],
                "summary": "List all namespaces or namespaces' ID",
                "parameters": [
                    {
                        "enum": [
                            "id"
                        ],
                        "type": "string",
                        "description": "Option",
                        "name": "option",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Different return structures by the given option param",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "[DEFAULT]": {
                                            "$ref": "#/definitions/common.RestGetAllNsResponse"
                                        },
                                        "[ID]": {
                                            "$ref": "#/definitions/common.IdList"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "post": {
                "description": "Create namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Namespace] Namespace management"
                ],
                "summary": "Create namespace",
                "parameters": [
                    {
                        "description": "Details for a new namespace",
                        "name": "nsReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.NsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.NsInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all namespaces",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Namespace] Namespace management"
                ],
                "summary": "Delete all namespaces",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            }
        },
        "/ns/{nsId}": {
            "get": {
                "description": "Get namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Namespace] Namespace management"
                ],
                "summary": "Get namespace",
                "parameters": [
                    {
                        "type": "string",
                        "default": "ns01",
                        "description": "Namespace ID",
                        "name": "nsId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.NsInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "put": {
                "description": "Update namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Namespace] Namespace management"
                ],
                "summary": "Update namespace",
                "parameters": [
                    {
                        "type": "string",
                        "default": "ns01",
                        "description": "Namespace ID",
                        "name": "nsId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Details to update existing namespace",
                        "name": "namespace",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.NsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.NsInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete namespace",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "[Namespace] Namespace management"
                ],
                "summary": "Delete namespace",
                "parameters": [
                    {
                        "type": "string",
                        "default": "ns01",
                        "description": "Namespace ID",
                        "name": "nsId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.SimpleMsg"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.ConfigInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string",
                    "example": "SPIDER_REST_URL"
                },
                "name": {
                    "type": "string",
                    "example": "SPIDER_REST_URL"
                },
                "value": {
                    "type": "string",
                    "example": "http://localhost:1024/spider"
                }
            }
        },
        "common.ConfigReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "SPIDER_REST_URL"
                },
                "value": {
                    "type": "string",
                    "example": "http://localhost:1024/spider"
                }
            }
        },
        "common.IdList": {
            "type": "object",
            "properties": {
                "output": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "common.JSONResult": {
            "type": "object"
        },
        "common.NsInfo": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Description for this namespace"
                },
                "id": {
                    "type": "string",
                    "example": "ns01"
                },
                "name": {
                    "type": "string",
                    "example": "ns01"
                }
            }
        },
        "common.NsReq": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Description for this namespace"
                },
                "name": {
                    "type": "string",
                    "example": "ns01"
                }
            }
        },
        "common.RestGetAllConfigResponse": {
            "type": "object",
            "properties": {
                "config": {
                    "description": "Name string     ` + "`" + `json:\"name\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.ConfigInfo"
                    }
                }
            }
        },
        "common.RestGetAllNsResponse": {
            "type": "object",
            "properties": {
                "ns": {
                    "description": "Name string     ` + "`" + `json:\"name\"` + "`" + `",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.NsInfo"
                    }
                }
            }
        },
        "common.SimpleMsg": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Any message"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "latest",
	Host:             "",
	BasePath:         "/beetle",
	Schemes:          []string{},
	Title:            "CM-Beetle REST API",
	Description:      "CM-Beetle REST API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
