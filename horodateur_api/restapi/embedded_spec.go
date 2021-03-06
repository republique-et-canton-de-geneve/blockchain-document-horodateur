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
  "swagger": "2.0",
  "info": {
    "description": "RCG horodatage est un service qui permet l'horodatage numérique via\nsur la blockchain Ethereum.\nLe principe est d'envoyer des fichiers qui sont ensuite passer dans\nune fonction hachage SHA3-256. Les « hash » sont ensuite intégrés\ndans un arbre de Merkle dont la racine est inséré dans une\ntransaction blockchain, l'(es) adresse(s) signant la transaction\nidentifie le Registre du Commerce, c'est une information qui doit\nêtre publique.\n",
    "title": "RCG horodatage",
    "version": "0.1.0"
  },
  "paths": {
    "/horodatage": {
      "get": {
        "description": "Liste les fichiers horodatés\n",
        "summary": "Liste les fichiers horodatés",
        "operationId": "listtimestamped",
        "responses": {
          "200": {
            "description": "Liste des fichiers qui ont été horodaté\n",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ReceiptFile"
              }
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/recu": {
      "get": {
        "description": "Retourne le fichier avec le reçu associé au hash fourni\n",
        "produces": [
          "application/octet-stream",
          "application/json"
        ],
        "summary": "Retourne le fichier avec le reçu",
        "operationId": "getreceipt",
        "parameters": [
          {
            "type": "string",
            "description": "Le hash identifiant un fichier",
            "name": "hash",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Langue du reçu",
            "name": "lang",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Fichier de reçu de l'horodatage certifié blockchain\n",
            "schema": {
              "type": "file"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Supprimer les reçus de la base de donnée\n",
        "summary": "Supprime les reçus de la base de donnée",
        "operationId": "delreceipts",
        "parameters": [
          {
            "description": "Liste des hash à supprimer",
            "name": "hashes",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/saml": {
      "get": {
        "description": "null",
        "summary": "null",
        "operationId": "configureSAML",
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/sonde": {
      "get": {
        "description": "Retourne quatres flag correspondant respectivement à la bonne connexion à un noeud Infura, la balance supérieure à 1 ETH, supérieure à 0,1 ETH et le bon fonctionnement d'une requête vers la base de données\n",
        "summary": "Retourne quatres flag correspondant respectivement à la bonne connexion à un noeud Infura, la balance supérieure à 1 ETH, supérieure à 0,1 ETH et le bon fonctionnement d'une requête vers la base de données.",
        "operationId": "monitoring",
        "responses": {
          "200": {
            "description": "Tout est en ordre et fonctionne correctement.\n",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Sonde"
              }
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
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
    "ReceiptFile": {
      "type": "object",
      "properties": {
        "date": {
          "type": "integer"
        },
        "filename": {
          "type": "string"
        },
        "hash": {
          "type": "string"
        },
        "horodatingaddress": {
          "type": "string"
        },
        "transactionhash": {
          "type": "string"
        }
      }
    },
    "Sonde": {
      "type": "object",
      "properties": {
        "balanceErrorThresholdExceeded": {
          "type": "boolean"
        },
        "balanceWarningThresholdExceeded": {
          "type": "boolean"
        },
        "ethereumActive": {
          "type": "boolean"
        },
        "persistenceActive": {
          "type": "boolean"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "RCG horodatage est un service qui permet l'horodatage numérique via\nsur la blockchain Ethereum.\nLe principe est d'envoyer des fichiers qui sont ensuite passer dans\nune fonction hachage SHA3-256. Les « hash » sont ensuite intégrés\ndans un arbre de Merkle dont la racine est inséré dans une\ntransaction blockchain, l'(es) adresse(s) signant la transaction\nidentifie le Registre du Commerce, c'est une information qui doit\nêtre publique.\n",
    "title": "RCG horodatage",
    "version": "0.1.0"
  },
  "paths": {
    "/horodatage": {
      "get": {
        "description": "Liste les fichiers horodatés\n",
        "summary": "Liste les fichiers horodatés",
        "operationId": "listtimestamped",
        "responses": {
          "200": {
            "description": "Liste des fichiers qui ont été horodaté\n",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/ReceiptFile"
              }
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/recu": {
      "get": {
        "description": "Retourne le fichier avec le reçu associé au hash fourni\n",
        "produces": [
          "application/octet-stream",
          "application/json"
        ],
        "summary": "Retourne le fichier avec le reçu",
        "operationId": "getreceipt",
        "parameters": [
          {
            "type": "string",
            "description": "Le hash identifiant un fichier",
            "name": "hash",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Langue du reçu",
            "name": "lang",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Fichier de reçu de l'horodatage certifié blockchain\n",
            "schema": {
              "type": "file"
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "description": "Supprimer les reçus de la base de donnée\n",
        "summary": "Supprime les reçus de la base de donnée",
        "operationId": "delreceipts",
        "parameters": [
          {
            "description": "Liste des hash à supprimer",
            "name": "hashes",
            "in": "body",
            "required": true,
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        ],
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/saml": {
      "get": {
        "description": "null",
        "summary": "null",
        "operationId": "configureSAML",
        "responses": {
          "200": {},
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/sonde": {
      "get": {
        "description": "Retourne quatres flag correspondant respectivement à la bonne connexion à un noeud Infura, la balance supérieure à 1 ETH, supérieure à 0,1 ETH et le bon fonctionnement d'une requête vers la base de données\n",
        "summary": "Retourne quatres flag correspondant respectivement à la bonne connexion à un noeud Infura, la balance supérieure à 1 ETH, supérieure à 0,1 ETH et le bon fonctionnement d'une requête vers la base de données.",
        "operationId": "monitoring",
        "responses": {
          "200": {
            "description": "Tout est en ordre et fonctionne correctement.\n",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Sonde"
              }
            }
          },
          "default": {
            "description": "Internal error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
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
    "ReceiptFile": {
      "type": "object",
      "properties": {
        "date": {
          "type": "integer"
        },
        "filename": {
          "type": "string"
        },
        "hash": {
          "type": "string"
        },
        "horodatingaddress": {
          "type": "string"
        },
        "transactionhash": {
          "type": "string"
        }
      }
    },
    "Sonde": {
      "type": "object",
      "properties": {
        "balanceErrorThresholdExceeded": {
          "type": "boolean"
        },
        "balanceWarningThresholdExceeded": {
          "type": "boolean"
        },
        "ethereumActive": {
          "type": "boolean"
        },
        "persistenceActive": {
          "type": "boolean"
        }
      }
    }
  }
}`))
}
