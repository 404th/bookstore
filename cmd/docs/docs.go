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
        "termsOfService": "http://terms_of_service_is_here.uzb",
        "contact": {
            "name": "API Support",
            "url": "http://t.me/four0fourth",
            "email": "umarov.doniyor.2001@gmail.com"
        },
        "license": {
            "name": "Incredible Certificate of Apache Co. 23923.0",
            "url": "http://www.incredible_certificate_of_apache_co.organisation/licenses/LICENSE-32423423423.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authors": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "get all authors",
                "operationId": "get_all_authors_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search query",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "has no relation with others",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Create an author",
                "operationId": "create_author_id",
                "parameters": [
                    {
                        "description": "author body",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateAuthor"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/authors/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "get author by ID",
                "operationId": "get_author_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "author id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Author"
                ],
                "summary": "Update Author",
                "operationId": "update_author_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "author id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "author update model",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateAuthor"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Author"
                ],
                "summary": "delete an author by id",
                "operationId": "delete_author_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "author id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/book_category": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Get all book categories",
                "operationId": "get_all_book_categories_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search query",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "has no relation with others",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Create a book category",
                "operationId": "create_book_category_id",
                "parameters": [
                    {
                        "description": "author body",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBookCategory"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/book_category/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Get book category by ID",
                "operationId": "get_book_category_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "BookCategory"
                ],
                "summary": "Update book category",
                "operationId": "update_author_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "book category update model",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBookCategory"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "BookCategory"
                ],
                "summary": "delete an book category by id",
                "operationId": "delete_book_category_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/books": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get all books",
                "operationId": "get_all_books_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "search",
                        "name": "search",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "has no relation with others",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Create a book",
                "operationId": "create_book_id",
                "parameters": [
                    {
                        "description": "book body",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBook"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        },
        "/books/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Get book by ID",
                "operationId": "get_book_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book category id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description of the RESPONSE",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Book"
                ],
                "summary": "Update book",
                "operationId": "update_book_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "book update model",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateBook"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Book"
                ],
                "summary": "delete an book by id",
                "operationId": "delete_book_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "book id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Description",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "400": {
                        "description": "Some bad request",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/models.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateAuthor": {
            "type": "object",
            "required": [
                "firstname",
                "secondname"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 22
                },
                "email": {
                    "type": "string",
                    "example": "created email"
                },
                "firstname": {
                    "type": "string",
                    "example": "Created John"
                },
                "secondname": {
                    "type": "string",
                    "example": "created secondname"
                }
            }
        },
        "models.CreateBook": {
            "type": "object",
            "required": [
                "author_id",
                "category_id",
                "name",
                "price"
            ],
            "properties": {
                "author_id": {
                    "type": "string",
                    "example": "author_id"
                },
                "category_id": {
                    "type": "string",
                    "example": "qwerty123"
                },
                "definition": {
                    "type": "string",
                    "example": "created very poor book"
                },
                "name": {
                    "type": "string",
                    "example": "Start with why"
                },
                "price": {
                    "type": "number",
                    "example": 18.99
                }
            }
        },
        "models.CreateBookCategory": {
            "type": "object",
            "required": [
                "category_name"
            ],
            "properties": {
                "category_name": {
                    "type": "string",
                    "example": "fiction"
                }
            }
        },
        "models.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.UpdateAuthor": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 26
                },
                "email": {
                    "type": "string",
                    "example": "updatedexample@example.com"
                },
                "firstname": {
                    "type": "string",
                    "example": "Updated John"
                },
                "secondname": {
                    "type": "string",
                    "example": "Updated Doe"
                }
            }
        },
        "models.UpdateBook": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "string",
                    "example": "qwerty123"
                },
                "category_id": {
                    "type": "string",
                    "example": "qwerty123"
                },
                "definition": {
                    "type": "string",
                    "example": "updated very poor book"
                },
                "name": {
                    "type": "string",
                    "example": "start with why"
                },
                "price": {
                    "type": "string",
                    "example": "18.99"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UpdateBookCategory": {
            "type": "object",
            "required": [
                "category_name"
            ],
            "properties": {
                "category_name": {
                    "type": "string",
                    "example": "fiction"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Book Store",
	Description:      "Book Store beta application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
