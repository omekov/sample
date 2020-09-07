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
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/podcasts": {
            "get": {
                "description": "Get details of all podcasts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "podcasts"
                ],
                "summary": "Get details of all podcasts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Podcast"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new podcast with the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "podcasts"
                ],
                "summary": "Create a new podcast",
                "parameters": [
                    {
                        "description": "Create podcast",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Podcast"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Podcast"
                        }
                    }
                }
            }
        },
        "/profile": {
            "post": {
                "description": "Profile customer the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sign"
                ],
                "summary": "Profile customer",
                "parameters": [
                    {
                        "description": "Profile customer",
                        "name": "signup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "description": "Sign auth client the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sign"
                ],
                "summary": "Sign auth",
                "parameters": [
                    {
                        "description": "SignIn auth",
                        "name": "signin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SignInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SignInput"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "description": "Sign Up new customer the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "sign"
                ],
                "summary": "Sign Up new customer",
                "parameters": [
                    {
                        "description": "SignUp customer",
                        "name": "signup",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Customer": {
            "type": "object",
            "properties": {
                "firstname": {
                    "type": "string",
                    "example": "Adam"
                },
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "registrationDate": {
                    "type": "string",
                    "example": "2020-09-09T21:21:46+00:00"
                },
                "releaseDate": {
                    "type": "string",
                    "example": "2020-09-09T22:21:46+00:00"
                },
                "username": {
                    "type": "string",
                    "example": "example@gmail.com"
                }
            }
        },
        "models.Podcast": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "example": "example@gmail.com"
                },
                "title": {
                    "type": "string",
                    "example": "title"
                }
            }
        },
        "models.SignInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "type": "string",
                    "example": "example@gmail.com"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
