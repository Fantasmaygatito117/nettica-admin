// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "produces": [
        "application/json"
    ],
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
        "/device": {
            "get": {
                "security": [
                    {
                        "apiKey": []
                    }
                ],
                "description": "Read all devices",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Read all devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth true \"X-API-KEY\" \"device-api-\u003capikey\u003e\"": []
                    },
                    {
                        "OAuth2": []
                    }
                ],
                "description": "Create a device",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Create a device",
                "parameters": [
                    {
                        "description": "model.Device",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Device"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {}
                    }
                }
            }
        },
        "/device/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth true \"X-API-KEY\" \"device-api-\u003capikey\u003e\"": []
                    },
                    {
                        "OAuth2": []
                    }
                ],
                "description": "Read a device",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "devices"
                ],
                "summary": "Read a device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Device": {
            "type": "object",
            "properties": {
                "accountid": {
                    "type": "string"
                },
                "apiKey": {
                    "type": "string"
                },
                "apiid": {
                    "type": "string"
                },
                "appdata": {
                    "type": "string"
                },
                "arch": {
                    "type": "string"
                },
                "authdomain": {
                    "type": "string"
                },
                "checkInterval": {
                    "type": "integer"
                },
                "clientid": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "debug": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                },
                "enable": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "lastSeen": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "os": {
                    "type": "string"
                },
                "platform": {
                    "type": "string"
                },
                "quiet": {
                    "type": "boolean"
                },
                "server": {
                    "type": "string"
                },
                "serviceApiKey": {
                    "type": "string"
                },
                "serviceGroup": {
                    "type": "string"
                },
                "sourceAddress": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                },
                "updatedBy": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                },
                "vpns": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.VPN"
                    }
                }
            }
        },
        "model.Settings": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "allowedIPs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "dns": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "enableDns": {
                    "type": "boolean"
                },
                "endpoint": {
                    "type": "string"
                },
                "listenPort": {
                    "type": "integer"
                },
                "mtu": {
                    "type": "integer"
                },
                "persistentKeepalive": {
                    "type": "integer"
                },
                "postDown": {
                    "type": "string"
                },
                "postUp": {
                    "type": "string"
                },
                "preDown": {
                    "type": "string"
                },
                "preUp": {
                    "type": "string"
                },
                "presharedKey": {
                    "type": "string"
                },
                "privateKey": {
                    "type": "string"
                },
                "publicKey": {
                    "type": "string"
                },
                "subnetRouting": {
                    "type": "boolean"
                },
                "table": {
                    "type": "string"
                },
                "upnp": {
                    "type": "boolean"
                }
            }
        },
        "model.VPN": {
            "type": "object",
            "properties": {
                "accountid": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "createdBy": {
                    "type": "string"
                },
                "current": {
                    "$ref": "#/definitions/model.Settings"
                },
                "default": {
                    "$ref": "#/definitions/model.Settings"
                },
                "deviceid": {
                    "type": "string"
                },
                "enable": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "netName": {
                    "type": "string"
                },
                "netid": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "type": "string"
                },
                "updated": {
                    "type": "string"
                },
                "updatedBy": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "apiKey": {
            "type": "apiKey",
            "name": "X-API-KEY",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "my.nettica.com",
	BasePath:         "/api/v1.0",
	Schemes:          []string{"https"},
	Title:            "Nettica API",
	Description:      "Nettica API documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
