// Code generated by swaggo/swag. DO NOT EDIT
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
        "/service/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取服务列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理接口"
                ],
                "summary": "获取服务列表",
                "operationId": "/service/list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "关键词",
                        "name": "info",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "页数",
                        "name": "page_no",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "每页条数",
                        "name": "page_size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/public.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schemas.ServiceListOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_add_http": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "添加HTTP服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理接口"
                ],
                "summary": "添加HTTP服务",
                "operationId": "/service/service_add_http",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ServiceAddHTTPInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/public.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/service_update_http": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "修改HTTP服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理接口"
                ],
                "summary": "修改HTTP服务",
                "operationId": "/service/service_update_http",
                "parameters": [
                    {
                        "description": "body",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ServiceUpdateHTTPInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/public.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/service/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "获取服务信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理接口"
                ],
                "summary": "获取服务信息",
                "operationId": "/service/{id}",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/public.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.ServiceDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "删除服务",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "服务管理接口"
                ],
                "summary": "删除服务",
                "operationId": "/service/{id}",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/public.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
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
        "models.HttpRule": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "need_websocket": {
                    "type": "integer"
                },
                "rule": {
                    "type": "string"
                },
                "service_id": {
                    "type": "integer"
                },
                "url_rewrite": {
                    "type": "string"
                }
            }
        },
        "models.ServiceDetail": {
            "type": "object",
            "properties": {
                "http_rule": {
                    "$ref": "#/definitions/models.HttpRule"
                },
                "info": {
                    "$ref": "#/definitions/models.ServiceInfo"
                }
            }
        },
        "models.ServiceInfo": {
            "type": "object",
            "properties": {
                "create_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_delete": {
                    "type": "integer"
                },
                "service_desc": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "public.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.ServiceAddHTTPInput": {
            "type": "object",
            "required": [
                "rule",
                "service_desc",
                "service_name"
            ],
            "properties": {
                "need_websocket": {
                    "description": "是否支持websocket",
                    "type": "integer",
                    "maximum": 1,
                    "minimum": 0,
                    "example": 0
                },
                "rule": {
                    "description": "域名或者前缀",
                    "type": "string",
                    "example": "类似/xxx/"
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1
                },
                "service_name": {
                    "description": "服务名",
                    "type": "string"
                },
                "url_rewrite": {
                    "description": "url重写功能",
                    "type": "string",
                    "example": "http://xx.xx.xx.xx:oo/"
                }
            }
        },
        "schemas.ServiceItemOutput": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "service_addr": {
                    "type": "string"
                },
                "service_desc": {
                    "type": "string"
                },
                "service_name": {
                    "type": "string"
                },
                "service_rewrite": {
                    "type": "string"
                }
            }
        },
        "schemas.ServiceListOutput": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.ServiceItemOutput"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "schemas.ServiceUpdateHTTPInput": {
            "type": "object",
            "required": [
                "id",
                "rule",
                "service_desc",
                "service_name"
            ],
            "properties": {
                "id": {
                    "description": "服务ID",
                    "type": "integer",
                    "minimum": 1,
                    "example": 62
                },
                "need_websocket": {
                    "description": "是否支持websocket",
                    "type": "integer",
                    "maximum": 1,
                    "minimum": 0
                },
                "rule": {
                    "description": "域名或者前缀 \t//启用strip_uri",
                    "type": "string",
                    "example": "/test_http_service_indb"
                },
                "service_desc": {
                    "description": "服务描述",
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 1,
                    "example": "test_http_service_indb"
                },
                "service_name": {
                    "description": "服务名",
                    "type": "string",
                    "example": "test_http_service_indb"
                },
                "url_rewrite": {
                    "description": "header转换",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
