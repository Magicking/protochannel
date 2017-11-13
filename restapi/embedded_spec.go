// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Protochannel API\n",
    "title": "Protochannel",
    "version": "0.1.0"
  },
  "paths": {
    "/commit": {
      "post": {
        "description": "Send and commit a valid signed ethereum message",
        "summary": "Commit a ethereum signed message to a channel",
        "operationId": "commit_to_channel",
        "parameters": [
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Message"
            }
          }
        ],
        "responses": {
          "200": {},
          "default": {
            "description": "Unexpected error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Return various information",
        "summary": "Various information on the micro service",
        "operationId": "status",
        "responses": {
          "200": {
            "description": "Information of the micro service",
            "schema": {
              "$ref": "#/definitions/Information"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Channel": {
      "description": "Channel information",
      "type": "object",
      "properties": {
        "channel": {
          "description": "Channel Identifier",
          "type": "string"
        },
        "nonce": {
          "description": "Channel nonce",
          "type": "integer"
        }
      }
    },
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "fields": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Information": {
      "type": "object",
      "properties": {
        "channels": {
          "description": "List of in-memory channel",
          "type": "array",
          "items": {
            "$ref": "#/definitions/Channel"
          }
        },
        "contract_address": {
          "description": "Channel address",
          "type": "string"
        }
      }
    },
    "Message": {
      "description": "Message to send to channel\n",
      "type": "object",
      "required": [
        "data"
      ],
      "properties": {
        "data": {
          "description": "Hex encoded payload(TODO TDB) information",
          "type": "string"
        }
      }
    }
  }
}`))
}