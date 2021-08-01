// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Greeting Server",
    "version": "1.0"
  },
  "host": "api.server.test",
  "basePath": "/",
  "paths": {
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains actual greeting msg",
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Greeting Server",
    "version": "1.0"
  },
  "host": "api.server.test",
  "basePath": "/",
  "paths": {
    "/hello": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "parameters": [
          {
            "type": "string",
            "description": "defaults to World if not given",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "returns a greeting",
            "schema": {
              "description": "contains actual greeting msg",
              "type": "string"
            }
          }
        }
      }
    }
  }
}`))
}