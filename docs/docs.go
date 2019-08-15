// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-15 16:48:47.484208708 +0800 CST m=+0.055643931

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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/account": {
            "get": {
                "description": "get Account libs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "account"
                ],
                "summary": "List Account",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.AccountPaginator"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            }
        },
        "/areas": {
            "get": {
                "description": "get areas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "area"
                ],
                "summary": "List Area",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/protocol.AreaPaginator"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "Create an area",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "area"
                ],
                "summary": "Create an area",
                "parameters": [
                    {
                        "description": "Add Area",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/protocol.AreaCreationParam"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/protocol.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            }
        },
        "/article": {
            "get": {
                "description": "get Articles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "article"
                ],
                "summary": "List Article",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.ArticlePaginator"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            }
        },
        "/question-libraries": {
            "get": {
                "description": "get question libs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question-libraries"
                ],
                "summary": "List question libs",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/protocol.QuestionLibraryPaginator"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            },
            "put": {
                "description": "add by json question lib",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question-libraries"
                ],
                "summary": "Add a question lib",
                "parameters": [
                    {
                        "description": "Add QuestionLibrary",
                        "name": "questionLib",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/controller.AddQuestionParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.QuestionLibrary"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            }
        },
        "/question-libraries/{id}": {
            "patch": {
                "description": "Update by json question lib",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "question-libraries"
                ],
                "summary": "Update a question lib",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "QuestionLib ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update account",
                        "name": "questionLib",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.QuestionLibrary"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.QuestionLibrary"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/protocol.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.AddQuestionParam": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "nestedParam": {
                    "type": "object",
                    "$ref": "#/definitions/controller.NestedParam"
                }
            }
        },
        "controller.NestedParam": {
            "type": "object",
            "required": [
                "nested2"
            ],
            "properties": {
                "nested1": {
                    "type": "string"
                },
                "nested2": {
                    "type": "integer"
                }
            }
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "pass": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "model.Area": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Area"
                    }
                },
                "code": {
                    "description": "区域编码",
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "description": "*gorm.Model",
                    "type": "integer"
                },
                "name": {
                    "description": "区域名称",
                    "type": "string"
                },
                "parent": {
                    "type": "object",
                    "$ref": "#/definitions/model.Area"
                },
                "parent_id": {
                    "description": "父级编号",
                    "type": "integer"
                },
                "parent_ids": {
                    "description": "所有父级编号",
                    "type": "string"
                },
                "type": {
                    "description": "区域类型（1：国家；2：省份、直辖市；3：地市；4：区县）",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "model.Article": {
            "type": "object",
            "properties": {
                "addDate": {
                    "description": "创建时间",
                    "type": "string"
                },
                "body": {
                    "description": "内容",
                    "type": "string"
                },
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "origin": {
                    "description": "用户名",
                    "type": "string"
                },
                "picture": {
                    "description": "图片",
                    "type": "string"
                },
                "postType": {
                    "description": "文章类型",
                    "type": "string"
                },
                "status": {
                    "description": "状态",
                    "type": "integer"
                },
                "subject": {
                    "description": "标题",
                    "type": "string"
                },
                "userID": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "visited": {
                    "description": "访问次数",
                    "type": "integer"
                }
            }
        },
        "model.QuestionLibrary": {
            "type": "object",
            "properties": {
                "difficulty": {
                    "type": "integer"
                },
                "inspectPoint": {
                    "type": "string"
                },
                "questionType": {
                    "type": "integer"
                },
                "securityLevel": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "subject": {
                    "type": "integer"
                },
                "techType": {
                    "type": "integer"
                }
            }
        },
        "protocol.AccountPaginator": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Account"
                    }
                },
                "total_page": {
                    "type": "integer"
                },
                "total_record": {
                    "type": "integer"
                }
            }
        },
        "protocol.AreaCreationParam": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "区域编码",
                    "type": "string"
                },
                "name": {
                    "description": "区域名称",
                    "type": "string"
                },
                "parent_id": {
                    "description": "父级编号",
                    "type": "integer"
                },
                "type": {
                    "description": "区域类型（1：国家；2：省份、直辖市；3：地市；4：区县）",
                    "type": "string"
                }
            }
        },
        "protocol.AreaPaginator": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Area"
                    }
                },
                "total_page": {
                    "type": "integer"
                },
                "total_record": {
                    "type": "integer"
                }
            }
        },
        "protocol.ArticlePaginator": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Article"
                    }
                },
                "total_page": {
                    "type": "integer"
                },
                "total_record": {
                    "type": "integer"
                }
            }
        },
        "protocol.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10001
                },
                "message": {
                    "type": "string",
                    "example": "Some Error"
                }
            }
        },
        "protocol.QuestionLibraryPaginator": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.QuestionLibrary"
                    }
                },
                "total_page": {
                    "type": "integer"
                },
                "total_record": {
                    "type": "integer"
                }
            }
        },
        "protocol.Success": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10000
                },
                "message": {
                    "type": "string",
                    "example": "Success"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "VCMS API",
	Description: "This is an api server celler server.",
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
