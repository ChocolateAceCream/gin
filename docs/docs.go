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
        "/v1/customer/query": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base/v1"
                ],
                "summary": "Get all user",
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\": { \"user\": { \"username\": \"asong\", \"nickname\": \"\", \"avatar\": \"\" }, \"token\": \"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VybmFtZSI6ImFzb25nIiwiZXhwIjoxNTk2OTAyMzEyLCJpc3MiOiJhc29uZyIsIm5iZiI6MTU5Njg5NDExMn0.uUS1TreZusX-hL3nKOSNYZIeZ_0BGrxWjKI6xdpdO40\", \"expiresAt\": 1596902312000 },,\"msg\":\"操作成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
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
