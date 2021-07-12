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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kube for Edge Management",
    "title": "Kube4EdgeManagement",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/k4e-management/v1",
  "paths": {
    "/control/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/control/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    }
  },
  "definitions": {
    "cpu": {
      "type": "object",
      "properties": {
        "architecture": {
          "type": "string"
        },
        "count": {
          "type": "integer"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "frequency": {
          "type": "number"
        },
        "model_name": {
          "type": "string"
        }
      }
    },
    "device-configuration": {
      "type": "object",
      "properties": {
        "heartbeat": {
          "$ref": "#/definitions/heartbeat-configuration"
        }
      }
    },
    "device-configuration-message": {
      "type": "object",
      "properties": {
        "configuration": {
          "$ref": "#/definitions/device-configuration"
        },
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "description": "List of workloads deployed to the device",
          "$ref": "#/definitions/workload-list"
        }
      }
    },
    "hardware-info": {
      "type": "object",
      "properties": {
        "cpu": {
          "$ref": "#/definitions/cpu"
        },
        "hostname": {
          "type": "string"
        },
        "memory": {
          "$ref": "#/definitions/memory"
        }
      }
    },
    "hardware-profile-configuration": {
      "type": "object",
      "properties": {
        "include": {
          "type": "boolean"
        },
        "scope": {
          "type": "string",
          "enum": [
            "full",
            "delta"
          ]
        }
      }
    },
    "heartbeat": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "enum": [
            "up",
            "degraded"
          ]
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/workload-status"
          }
        }
      }
    },
    "heartbeat-configuration": {
      "type": "object",
      "properties": {
        "hardware_profile": {
          "$ref": "#/definitions/hardware-profile-configuration"
        },
        "period_seconds": {
          "type": "integer"
        }
      }
    },
    "memory": {
      "type": "object",
      "properties": {
        "physical_bytes": {
          "type": "integer"
        },
        "usable_bytes": {
          "type": "integer"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "content": {
          "description": "Content"
        },
        "directive": {
          "type": "string"
        },
        "message_id": {
          "type": "string"
        },
        "metadata": {
          "type": "object"
        },
        "response_to": {
          "type": "string"
        },
        "sent": {
          "type": "string",
          "format": "date-time"
        },
        "type": {
          "type": "string",
          "enum": [
            "connection-status",
            "command",
            "event",
            "data"
          ]
        },
        "version": {
          "type": "integer"
        }
      }
    },
    "registration-info": {
      "type": "object",
      "properties": {
        "hardware": {
          "description": "Hardware information",
          "$ref": "#/definitions/hardware-info"
        },
        "os_image_id": {
          "type": "string"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "name": {
          "description": "Name of the workload",
          "type": "string"
        },
        "specification": {
          "type": "string"
        }
      }
    },
    "workload-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    },
    "workload-status": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "enum": [
            "deploying",
            "running",
            "crashed",
            "stopped"
          ]
        }
      }
    }
  },
  "securityDefinitions": {
    "agentAuth": {
      "type": "apiKey",
      "name": "X-Secret-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Device management",
      "name": "devices"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Kube for Edge Management",
    "title": "Kube4EdgeManagement",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/api/k4e-management/v1",
  "paths": {
    "/control/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/control/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostControlMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/in": {
      "get": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "GetDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    },
    "/data/{device_id}/out": {
      "post": {
        "tags": [
          "yggdrasil"
        ],
        "operationId": "PostDataMessageForDevice",
        "parameters": [
          {
            "type": "string",
            "description": "Device ID",
            "name": "device_id",
            "in": "path",
            "required": true
          },
          {
            "name": "message",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success"
          },
          "401": {
            "description": "Unauthorized"
          },
          "403": {
            "description": "Forbidden"
          },
          "404": {
            "description": "Error"
          },
          "500": {
            "description": "Error"
          }
        }
      }
    }
  },
  "definitions": {
    "cpu": {
      "type": "object",
      "properties": {
        "architecture": {
          "type": "string"
        },
        "count": {
          "type": "integer"
        },
        "flags": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "frequency": {
          "type": "number"
        },
        "model_name": {
          "type": "string"
        }
      }
    },
    "device-configuration": {
      "type": "object",
      "properties": {
        "heartbeat": {
          "$ref": "#/definitions/heartbeat-configuration"
        }
      }
    },
    "device-configuration-message": {
      "type": "object",
      "properties": {
        "configuration": {
          "$ref": "#/definitions/device-configuration"
        },
        "device_id": {
          "description": "Device identifier",
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "description": "List of workloads deployed to the device",
          "$ref": "#/definitions/workload-list"
        }
      }
    },
    "hardware-info": {
      "type": "object",
      "properties": {
        "cpu": {
          "$ref": "#/definitions/cpu"
        },
        "hostname": {
          "type": "string"
        },
        "memory": {
          "$ref": "#/definitions/memory"
        }
      }
    },
    "hardware-profile-configuration": {
      "type": "object",
      "properties": {
        "include": {
          "type": "boolean"
        },
        "scope": {
          "type": "string",
          "enum": [
            "full",
            "delta"
          ]
        }
      }
    },
    "heartbeat": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string",
          "enum": [
            "up",
            "degraded"
          ]
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "version": {
          "type": "string"
        },
        "workloads": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/workload-status"
          }
        }
      }
    },
    "heartbeat-configuration": {
      "type": "object",
      "properties": {
        "hardware_profile": {
          "$ref": "#/definitions/hardware-profile-configuration"
        },
        "period_seconds": {
          "type": "integer"
        }
      }
    },
    "memory": {
      "type": "object",
      "properties": {
        "physical_bytes": {
          "type": "integer"
        },
        "usable_bytes": {
          "type": "integer"
        }
      }
    },
    "message": {
      "type": "object",
      "properties": {
        "content": {
          "description": "Content"
        },
        "directive": {
          "type": "string"
        },
        "message_id": {
          "type": "string"
        },
        "metadata": {
          "type": "object"
        },
        "response_to": {
          "type": "string"
        },
        "sent": {
          "type": "string",
          "format": "date-time"
        },
        "type": {
          "type": "string",
          "enum": [
            "connection-status",
            "command",
            "event",
            "data"
          ]
        },
        "version": {
          "type": "integer"
        }
      }
    },
    "registration-info": {
      "type": "object",
      "properties": {
        "hardware": {
          "description": "Hardware information",
          "$ref": "#/definitions/hardware-info"
        },
        "os_image_id": {
          "type": "string"
        }
      }
    },
    "workload": {
      "type": "object",
      "properties": {
        "name": {
          "description": "Name of the workload",
          "type": "string"
        },
        "specification": {
          "type": "string"
        }
      }
    },
    "workload-list": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/workload"
      }
    },
    "workload-status": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "status": {
          "type": "string",
          "enum": [
            "deploying",
            "running",
            "crashed",
            "stopped"
          ]
        }
      }
    }
  },
  "securityDefinitions": {
    "agentAuth": {
      "type": "apiKey",
      "name": "X-Secret-Key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Device management",
      "name": "devices"
    }
  ]
}`))
}
