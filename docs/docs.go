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
        "contact": {
            "name": "BFlameSwift",
            "url": "https://github.com/BFlameSwift",
            "email": "bflaemswift@163.com"
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
        "/login": {
            "post": {
                "description": "根据用户邮箱和密码等生成token，并将token返回给用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户名，密码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/response.LoginQ"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\": true, \"message\": \"登录成功\", \"data\": \"model.User的所有信息\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/resource/upload": {
            "post": {
                "description": "上传一个文件测试",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "其他模块"
                ],
                "summary": "上传文件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "是否成功，返回信息",
                        "schema": {
                            "$ref": "#/definitions/response.CommonA"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.CommonA": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "response.LoginQ": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Demo API",
	Description:      "Demo API description",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
