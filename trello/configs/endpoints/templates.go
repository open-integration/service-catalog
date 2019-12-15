
// Code generated by go generate; DO NOT EDIT.
// using data from templates/endpoints
package endpoints

func TemplatesMap() map[string]string {
    templatesMap := make(map[string]string)

templatesMap["GetCards/arguments.json"] = `{
    "$id": "https://example.com/person.schema.json",
    "$schema": "http://json-schema.org/draft-07/schema#",
    "title": "Trello GetCards Task Arguments Schema",
    "type": "object",
    "properties": {
      "App": {
        "type": "string",
        "description": "Trello Application ID",
        "minLength": 1
      },
      "Token": {
        "type": "string",
        "description": "Trello Token",
        "minLength": 1
      },
      "Board": {
        "description": "Trello Board ID",
        "type": "string",
        "minLength": 1
      }
    },
    "required": [
        "App",
        "Token",
        "Board"
    ]
  }` 

templatesMap["GetCards/returns.json"] = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "array",
  "items": {
      "$ref": "#/definitions/Card"
  },
  "definitions": {
      "Card": {
          "type": "object",
          "additionalProperties": false,
          "properties": {
              "id": {
                  "type": "string"
              },
              "idShort": {
                  "type": "number"
              },
              "name": {
                  "type": "string"
              },
              "pos": {
                  "type": "number"
              },
              "email": {
                  "type": "string"
              },
              "shortLink": {
                  "type": "string"
              },
              "shortUrl": {
                  "type": "string",
                  "format": "uri",
                  "qt-uri-protocols": [
                      "https"
                  ]
              },
              "url": {
                  "type": "string",
                  "format": "uri",
                  "qt-uri-protocols": [
                      "https"
                  ]
              },
              "desc": {
                  "type": "string"
              },
              "due": {
                "anyOf": [
                    { "type": "string" },
                    { "type": "null" }
                  ]
              },
              "dueComplete": {
                  "type": "boolean"
              },
              "closed": {
                  "type": "boolean"
              },
              "subscribed": {
                  "type": "boolean"
              },
              "dateLastActivity": {
                  "type": "string",
                  "format": "date-time"
              },
              "Board": {
                  "type": "null"
              },
              "idBoard": {
                  "type": "string"
              },
              "List": {
                  "$ref": "#/definitions/List"
              },
              "idList": {
                  "type": "string"
              },
              "badges": {
                  "$ref": "#/definitions/Badges"
              },
              "idCheckLists": {
                  "type": "array",
                  "items": {}
              },
              "idAttachmentCover": {
                  "type": "string"
              },
              "manualCoverAttachment": {
                  "type": "boolean"
              },
              "idLabels": {
                  "type": "array",
                  "items": {
                      "type": "string"
                  }
              },
              "labels": {
                  "type": "array",
                  "items": {
                      "$ref": "#/definitions/Label"
                  }
              }
          },
          "required": [
              "Board",
              "List",
              "badges",
              "closed",
              "dateLastActivity",
              "desc",
              "due",
              "dueComplete",
              "email",
              "id",
              "idAttachmentCover",
              "idBoard",
              "idCheckLists",
              "idList",
              "idShort",
              "manualCoverAttachment",
              "name",
              "pos",
              "shortLink",
              "shortUrl",
              "subscribed",
              "url"
          ],
          "title": "Card"
      },
      "List": {
          "type": "object",
          "additionalProperties": false,
          "properties": {
              "id": {
                  "type": "string"
              },
              "name": {
                  "type": "string"
              },
              "idBoard": {
                  "type": "string"
              },
              "closed": {
                  "type": "boolean"
              },
              "pos": {
                  "type": "number"
              },
              "subscribed": {
                  "type": "boolean"
              }
          },
          "required": [
              "closed",
              "id",
              "idBoard",
              "name",
              "pos",
              "subscribed"
          ],
          "title": "List"
      },
      "Badges": {
          "type": "object",
          "properties": {
              "votes": {
                  "type": "number"
              },
              "viewingMemberVoted": {
                  "type": "boolean"
              },
              "subscribed": {
                  "type": "boolean"
              },
              "checkItems": {
                  "type": "number"
              },
              "checkItemsChecked": {
                  "type": "number"
              },
              "comments": {
                  "type": "number"
              },
              "attachments": {
                  "type": "number"
              },
              "description": {
                  "type": "boolean"
              }
          },
          "required": [
              "attachments",
              "checkItems",
              "checkItemsChecked",
              "comments",
              "description",
              "subscribed",
              "viewingMemberVoted",
              "votes"
          ],
          "title": "Badges"
      },
      "Label": {
          "type": "object",
          "additionalProperties": false,
          "properties": {
              "id": {
                  "type": "string"
              },
              "idBoard": {
                  "type": "string"
              },
              "name": {
                  "type": "string"
              },
              "color": {
                  "type": "string"
              },
              "uses": {
                  "type": "number"
              }
          },
          "required": [
              "color",
              "id",
              "idBoard",
              "name",
              "uses"
          ],
          "title": "Label"
      }
  }
}
` 

    return  templatesMap
}
