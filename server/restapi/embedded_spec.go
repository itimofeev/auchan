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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a server for collaborative shopping.\n",
    "title": "CityProject for Auchan",
    "version": "1.0.0"
  },
  "host": "aeuchan.swagger.io",
  "basePath": "/v1",
  "paths": {
    "/basket": {
      "get": {
        "tags": [
          "basket"
        ],
        "summary": "get all users baskets",
        "operationId": "getAllBaskets",
        "responses": {
          "200": {
            "description": "returns all available baskets",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Basket"
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "basket"
        ],
        "summary": "create basket",
        "operationId": "createBasket",
        "parameters": [
          {
            "name": "basket",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Basket"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    },
    "/basket/{basketId}/goods": {
      "get": {
        "tags": [
          "goods"
        ],
        "summary": "return all goods in basket",
        "operationId": "getAllGoodsInBasket",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns goods in basket",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Goods"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "goods"
        ],
        "summary": "change goods quantity in basket",
        "operationId": "addGoodsToBasket",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          },
          {
            "name": "goods",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "productId": {
                  "type": "string"
                },
                "quantity": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns current state of goods",
            "schema": {
              "$ref": "#/definitions/Goods"
            }
          }
        }
      }
    },
    "/basket/{basketId}/share": {
      "get": {
        "tags": [
          "share"
        ],
        "summary": "get all shares for basket",
        "operationId": "getAllSharesForBasket",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns created share",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Share"
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "share"
        ],
        "summary": "add user to share",
        "operationId": "addUserToShare",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          },
          {
            "name": "share",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns created share",
            "schema": {
              "$ref": "#/definitions/Share"
            }
          }
        }
      }
    },
    "/product": {
      "get": {
        "tags": [
          "product"
        ],
        "summary": "search for products",
        "operationId": "getProductsByParams",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Product"
              }
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "type": "string",
            "description": "The user email for login",
            "name": "email",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "The password for login in clear text",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string"
            }
          },
          "401": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/{email}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by email",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "email",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Basket": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Goods": {
      "type": "object",
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "price": {
          "type": "integer"
        },
        "product": {
          "$ref": "#/definitions/Product"
        },
        "quantity": {
          "type": "integer"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "imageUrl": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Share": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user"
    },
    {
      "description": "Operations about basket of goods",
      "name": "basket"
    },
    {
      "description": "Operations about products",
      "name": "product"
    },
    {
      "description": "Operations about goods (product added to the basket with price and quantity)",
      "name": "goods"
    },
    {
      "description": "Operations about share of basket with somebody",
      "name": "share"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a server for collaborative shopping.\n",
    "title": "CityProject for Auchan",
    "version": "1.0.0"
  },
  "host": "aeuchan.swagger.io",
  "basePath": "/v1",
  "paths": {
    "/basket": {
      "get": {
        "tags": [
          "basket"
        ],
        "summary": "get all users baskets",
        "operationId": "getAllBaskets",
        "responses": {
          "200": {
            "description": "returns all available baskets",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Basket"
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "basket"
        ],
        "summary": "create basket",
        "operationId": "createBasket",
        "parameters": [
          {
            "name": "basket",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/Basket"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    },
    "/basket/{basketId}/goods": {
      "get": {
        "tags": [
          "goods"
        ],
        "summary": "return all goods in basket",
        "operationId": "getAllGoodsInBasket",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns goods in basket",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Goods"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "goods"
        ],
        "summary": "change goods quantity in basket",
        "operationId": "addGoodsToBasket",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          },
          {
            "name": "goods",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "productId": {
                  "type": "string"
                },
                "quantity": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns current state of goods",
            "schema": {
              "$ref": "#/definitions/Goods"
            }
          }
        }
      }
    },
    "/basket/{basketId}/share": {
      "get": {
        "tags": [
          "share"
        ],
        "summary": "get all shares for basket",
        "operationId": "getAllSharesForBasket",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "returns created share",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Share"
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "share"
        ],
        "summary": "add user to share",
        "operationId": "addUserToShare",
        "parameters": [
          {
            "type": "string",
            "name": "basketId",
            "in": "path",
            "required": true
          },
          {
            "name": "share",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "returns created share",
            "schema": {
              "$ref": "#/definitions/Share"
            }
          }
        }
      }
    },
    "/product": {
      "get": {
        "tags": [
          "product"
        ],
        "summary": "search for products",
        "operationId": "getProductsByParams",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "name",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Product"
              }
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    },
    "/user/login": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "type": "string",
            "description": "The user email for login",
            "name": "email",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "The password for login in clear text",
            "name": "password",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "type": "string"
            }
          },
          "401": {
            "description": "Invalid username/password supplied"
          }
        }
      }
    },
    "/user/logout": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs out current logged in user session",
        "operationId": "logoutUser",
        "responses": {
          "default": {
            "description": "successful operation"
          }
        }
      }
    },
    "/user/{email}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Get user by email",
        "operationId": "getUserByName",
        "parameters": [
          {
            "type": "string",
            "description": "The name that needs to be fetched.",
            "name": "email",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "400": {
            "description": "Invalid username supplied"
          },
          "404": {
            "description": "User not found"
          }
        }
      }
    }
  },
  "definitions": {
    "Basket": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Goods": {
      "type": "object",
      "properties": {
        "completed": {
          "type": "boolean"
        },
        "id": {
          "type": "string"
        },
        "price": {
          "type": "integer"
        },
        "product": {
          "$ref": "#/definitions/Product"
        },
        "quantity": {
          "type": "integer"
        },
        "unit": {
          "type": "string"
        }
      }
    },
    "Product": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "imageUrl": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "Share": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "userId": {
          "type": "string"
        }
      }
    },
    "User": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        }
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "Operations about user",
      "name": "user"
    },
    {
      "description": "Operations about basket of goods",
      "name": "basket"
    },
    {
      "description": "Operations about products",
      "name": "product"
    },
    {
      "description": "Operations about goods (product added to the basket with price and quantity)",
      "name": "goods"
    },
    {
      "description": "Operations about share of basket with somebody",
      "name": "share"
    }
  ]
}`))
}