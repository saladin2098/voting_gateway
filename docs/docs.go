// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/candidate/all": {
            "get": {
                "description": "Gets all candidates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidate"
                ],
                "summary": "gets all Candidate",
                "operationId": "get_all_candidate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllCandidate"
                        }
                    },
                    "500": {
                        "description": "Failed to get all candidates",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/candidate/by-party": {
            "get": {
                "description": "Gets all candidates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidate"
                ],
                "summary": "gets Candidate by the party id",
                "operationId": "get_candidates_by_party",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllCandidate"
                        }
                    },
                    "500": {
                        "description": "Failed to get candidates by party",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/candidate/create": {
            "post": {
                "description": "Creates candidate by reading from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidate"
                ],
                "summary": "Creates Candidate",
                "operationId": "create_candidate",
                "parameters": [
                    {
                        "description": "candiate body data",
                        "name": "Candiate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.CandidateCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Candiate created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create candidate",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/candidate/update": {
            "put": {
                "description": "updates candidate by reading from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidate"
                ],
                "summary": "updates Candidate",
                "operationId": "updates_candidate",
                "parameters": [
                    {
                        "description": "candiate data",
                        "name": "candiate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.Candidate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Candiate updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create candidate",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/candidate/{id}": {
            "get": {
                "description": "Gets candidate by reading id from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidate"
                ],
                "summary": "gets Candidate by id",
                "operationId": "get_candidate_by_id",
                "parameters": [
                    {
                        "type": "string",
                        "name": "Id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.Candidate"
                        }
                    },
                    "500": {
                        "description": "Failed to get candidate by id",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes candidate by reading id from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Candidate"
                ],
                "summary": "deletes Candidate",
                "operationId": "delete_candidate",
                "parameters": [
                    {
                        "type": "string",
                        "description": "candiate id data",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Candiate deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create candidate",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/election/all": {
            "get": {
                "description": "Gets all Elections",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Election"
                ],
                "summary": "Gets all Elections",
                "operationId": "get_all_elections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllElection"
                        }
                    },
                    "500": {
                        "description": "Failed to Get all elections",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/election/by-date": {
            "get": {
                "description": "Getss Elections by date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Election"
                ],
                "summary": "Getss Elections by date",
                "operationId": "get_elections_by_date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "election date",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllElection"
                        }
                    },
                    "500": {
                        "description": "Failed to Get elections data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/election/create": {
            "post": {
                "description": "Ccreate election by reading from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Election"
                ],
                "summary": "Creates Election",
                "operationId": "create_election",
                "parameters": [
                    {
                        "description": "election body data",
                        "name": "Election",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.Election"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.Void"
                        }
                    },
                    "500": {
                        "description": "Failed to get product by id",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/election/update": {
            "put": {
                "description": "Updates election by reading election from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Election"
                ],
                "summary": "Update Election",
                "operationId": "update_election",
                "parameters": [
                    {
                        "description": "election data",
                        "name": "election",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.Election"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.Void"
                        }
                    },
                    "500": {
                        "description": "Failed to update election data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/election/{id}": {
            "get": {
                "description": "gets election by reading id  from path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Election"
                ],
                "summary": "Gets Election by id",
                "operationId": "get_election_by_id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "election ID data",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.Election"
                        }
                    },
                    "500": {
                        "description": "Failed to get product by id",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes election by reading id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Election"
                ],
                "summary": "Deletes Election",
                "operationId": "delete_election",
                "parameters": [
                    {
                        "type": "string",
                        "description": "election ID data",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to delete election data",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public-vote/all": {
            "get": {
                "description": "Gets all Public Votes without any filter",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public-Votes"
                ],
                "summary": "Gets all Public Votes",
                "operationId": "get_all_public_votes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.GetAllPublicVote"
                        }
                    },
                    "500": {
                        "description": "Failed to get public votes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public-vote/create": {
            "post": {
                "description": "Create Public Vote by reading from body",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public-Votes"
                ],
                "summary": "Creates Public Vote",
                "operationId": "create_public_vote",
                "parameters": [
                    {
                        "description": "Public Vote body data",
                        "name": "PublicVote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.PublicVoteCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to create public vote",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public-vote/delete": {
            "delete": {
                "description": "Deletes a Public Vote by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public-Votes"
                ],
                "summary": "Deletes a Public Vote",
                "operationId": "delete_public_votes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Public Vote ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to delete public vote",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public-vote/find-winner": {
            "get": {
                "description": "finds the cnadidate withmost botes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public-Votes"
                ],
                "summary": "find the winner",
                "operationId": "fin_the_winner",
                "parameters": [
                    {
                        "type": "string",
                        "description": "election ID",
                        "name": "election",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.Winner"
                        }
                    },
                    "500": {
                        "description": "Failed to get public vote by ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public-vote/update": {
            "put": {
                "description": "Updates a Public Vote by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public-Votes"
                ],
                "summary": "Updates a Public Vote",
                "operationId": "update_public_votes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Public Vote ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Election ID",
                        "name": "election",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Public ID",
                        "name": "public",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to update public vote",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/public-vote/{id}": {
            "get": {
                "description": "Gets a Public Vote by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Public-Votes"
                ],
                "summary": "Gets a Public Vote by ID",
                "operationId": "get_by_id_public_votes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Public Vote ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.PublicVote"
                        }
                    },
                    "500": {
                        "description": "Failed to get public vote by ID",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.Candidate": {
            "type": "object",
            "properties": {
                "election": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "party": {
                    "type": "string"
                },
                "public": {
                    "type": "string"
                }
            }
        },
        "genprotos.CandidateCreate": {
            "type": "object",
            "properties": {
                "ElectionId": {
                    "type": "string"
                },
                "PartyId": {
                    "type": "string"
                },
                "PublicId": {
                    "type": "string"
                }
            }
        },
        "genprotos.Election": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "genprotos.GetAllCandidate": {
            "type": "object",
            "properties": {
                "Candidates": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.Candidate"
                    }
                }
            }
        },
        "genprotos.GetAllElection": {
            "type": "object",
            "properties": {
                "elections": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.Election"
                    }
                }
            }
        },
        "genprotos.GetAllPublicVote": {
            "type": "object",
            "properties": {
                "publicVotes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.PublicVote"
                    }
                }
            }
        },
        "genprotos.PublicVote": {
            "type": "object",
            "properties": {
                "CandidateId": {
                    "type": "string"
                },
                "ElectionId": {
                    "type": "string"
                },
                "Id": {
                    "type": "string"
                },
                "PublicId": {
                    "type": "string"
                }
            }
        },
        "genprotos.PublicVoteCreate": {
            "type": "object",
            "properties": {
                "CandidateId": {
                    "type": "string"
                },
                "ElectionId": {
                    "type": "string"
                },
                "PublicId": {
                    "type": "string"
                }
            }
        },
        "genprotos.Void": {
            "type": "object"
        },
        "genprotos.Winner": {
            "type": "object",
            "properties": {
                "CandidateId": {
                    "type": "string"
                },
                "ElectionId": {
                    "type": "string"
                },
                "Votes": {
                    "type": "integer"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
