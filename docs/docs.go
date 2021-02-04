// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "todo-list-hateoas.herokuapp.com",
        "contact": {
            "name": "Vinícius Boscardin",
            "email": "boscardinvinicius@gmail.com"
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
        "/item": {
            "get": {
                "description": "Search tasks by acronym",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Search tasks by acronym",
                "parameters": [
                    {
                        "type": "string",
                        "description": "vin",
                        "name": "sigla",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "allOf": [
                                    {
                                        "$ref": "#/definitions/domain.Item"
                                    },
                                    {
                                        "type": "object",
                                        "properties": {
                                            "links": {
                                                "type": "array",
                                                "items": {
                                                    "$ref": "#/definitions/domain.Link"
                                                }
                                            }
                                        }
                                    }
                                ]
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Include tasks into database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Include tasks into database",
                "parameters": [
                    {
                        "description": "item",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Item"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Item"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "links": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/domain.Link"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/item/{id}": {
            "get": {
                "description": "Search tasks by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Search tasks by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Item"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "links": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/domain.Link"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Change tasks into database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Change tasks into database",
                "parameters": [
                    {
                        "description": "item",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Item"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Item"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "links": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/domain.Link"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete tasks",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Delete tasks",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/domain.Item"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "links": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/domain.Link"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Item": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "2021-02-02"
                },
                "descricao": {
                    "type": "string",
                    "example": "Descrição da tarefa 1"
                },
                "nome": {
                    "type": "string",
                    "example": "Tarefa 1"
                },
                "sigla": {
                    "type": "string",
                    "maxLength": 3,
                    "example": "vin"
                }
            }
        },
        "domain.Link": {
            "type": "object",
            "properties": {
                "href": {
                    "type": "string",
                    "example": "http(s)://\u003cDOMAIN_OR_IP\u003e/item/{id}"
                },
                "method": {
                    "type": "string",
                    "example": "GET"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "2021.2.1.0",
	Host:        "todo-list-hateoas.herokuapp.com",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Clean archtecture and Level 3 of REST",
	Description: "An application of studies on the implementation of clean architecture with golang with a plus of REST level 3 implementations",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
