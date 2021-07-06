package v2

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"

	"github.com/stretchr/testify/require"
)

var doc = []byte(`{
  "asyncapi": "2.0.0",
  "x-test": "foobar",
  "channels": {
    "smartylighting/streetlights/1/0/action/{streetlightId}/dim": {
      "parameters": {
        "streetlightId": {
          "description": "The ID of the streetlight.",
          "schema": {
            "type": "string"
          }
        }
      },
      "subscribe": {
        "message": {
          "name": "dimLight",
          "payload": {
            "properties": {
              "percentage": {
                "description": "Percentage to which the light should be dimmed to.",
                "maximum": 100,
                "minimum": 0,
                "type": "integer"
              },
              "sentAt": {
                "description": "Date and time when the message was sent.",
                "format": "date-time",
                "type": "string"
              }
            },
            "type": "object"
          },
          "summary": "Command a particular streetlight to dim the lights.",
          "title": "Dim light",
          "traits": [
            {
              "headers": {
                "properties": {
                  "my-app-header": {
                    "maximum": 100,
                    "minimum": 0,
                    "type": "integer"
                  }
                },
                "type": "object"
              }
            }
          ]
        },
        "operationId": "dimLight",
        "traits": [
          {
            "bindings": {
              "mqtt": {
                "qos": 1
              }
            }
          }
        ]
      }
    },
    "smartylighting/streetlights/1/0/action/{streetlightId}/turn/off": {
      "parameters": {
        "streetlightId": {
          "description": "The ID of the streetlight.",
          "schema": {
            "type": "string"
          }
        }
      },
      "subscribe": {
        "message": {
          "name": "turnOnOff",
          "payload": {
            "properties": {
              "command": {
                "description": "Whether to turn on or off the light.",
                "enum": [
                  "on",
                  "off"
                ],
                "type": "string"
              },
              "sentAt": {
                "description": "Date and time when the message was sent.",
                "format": "date-time",
                "type": "string"
              }
            },
            "type": "object"
          },
          "summary": "Command a particular streetlight to turn the lights on or off.",
          "title": "Turn on/off",
          "traits": [
            {
              "headers": {
                "properties": {
                  "my-app-header": {
                    "maximum": 100,
                    "minimum": 0,
                    "type": "integer"
                  }
                },
                "type": "object"
              }
            }
          ]
        },
        "operationId": "turnOff",
        "traits": [
          {
            "bindings": {
              "mqtt": {
                "qos": 1
              }
            }
          }
        ]
      }
    },
    "smartylighting/streetlights/1/0/action/{streetlightId}/turn/on": {
      "parameters": {
        "streetlightId": {
          "description": "The ID of the streetlight.",
          "schema": {
            "type": "string"
          }
        }
      },
      "subscribe": {
        "message": {
          "name": "turnOnOff",
          "payload": {
            "properties": {
              "command": {
                "description": "Whether to turn on or off the light.",
                "enum": [
                  "on",
                  "off"
                ],
                "type": "string"
              },
              "sentAt": {
                "description": "Date and time when the message was sent.",
                "format": "date-time",
                "type": "string"
              }
            },
            "type": "object"
          },
          "summary": "Command a particular streetlight to turn the lights on or off.",
          "title": "Turn on/off",
          "traits": [
            {
              "headers": {
                "properties": {
                  "my-app-header": {
                    "maximum": 100,
                    "minimum": 0,
                    "type": "integer"
                  }
                },
                "type": "object"
              }
            }
          ]
        },
        "operationId": "turnOn",
        "traits": [
          {
            "bindings": {
              "mqtt": {
                "qos": 1
              }
            }
          }
        ]
      }
    },
    "smartylighting/streetlights/1/0/event/{streetlightId}/lighting/measured": {
      "description": "The topic on which measured values may be produced and consumed.",
      "parameters": {
        "streetlightId": {
          "description": "The ID of the streetlight.",
          "schema": {
            "type": "string"
          }
        }
      },
      "publish": {
        "message": {
          "contentType": "application/json",
          "name": "lightMeasured",
          "payload": {
            "properties": {
              "lumens": {
                "description": "Light intensity measured in lumens.",
                "minimum": 0,
                "type": "integer"
              },
              "sentAt": {
                "description": "Date and time when the message was sent.",
                "format": "date-time",
                "type": "string"
              }
            },
            "type": "object"
          },
          "summary": "Inform about environmental lighting conditions of a particular streetlight.",
          "title": "Light measured",
          "traits": [
            {
              "headers": {
                "properties": {
                  "my-app-header": {
                    "maximum": 100,
                    "minimum": 0,
                    "type": "integer"
                  }
                },
                "type": "object"
              }
            }
          ]
        },
        "operationId": "receiveLightMeasurement",
        "summary": "Inform about environmental lighting conditions of a particular streetlight.",
        "traits": [
          {
            "bindings": {
              "mqtt": {
                "qos": 1
              }
            }
          }
        ]
      }
    }
  },
  "components": {
    "messageTraits": {
      "commonHeaders": {
        "headers": {
          "properties": {
            "my-app-header": {
              "maximum": 100,
              "minimum": 0,
              "type": "integer"
            }
          },
          "type": "object"
        }
      }
    },
    "messages": {
      "dimLight": {
        "name": "dimLight",
        "payload": {
          "properties": {
            "percentage": {
              "description": "Percentage to which the light should be dimmed to.",
              "maximum": 100,
              "minimum": 0,
              "type": "integer"
            },
            "sentAt": {
              "description": "Date and time when the message was sent.",
              "format": "date-time",
              "type": "string"
            }
          },
          "type": "object"
        },
        "summary": "Command a particular streetlight to dim the lights.",
        "title": "Dim light",
        "traits": [
          {
            "headers": {
              "properties": {
                "my-app-header": {
                  "maximum": 100,
                  "minimum": 0,
                  "type": "integer"
                }
              },
              "type": "object"
            }
          }
        ]
      },
      "lightMeasured": {
        "contentType": "application/json",
        "name": "lightMeasured",
        "payload": {
          "properties": {
            "lumens": {
              "description": "Light intensity measured in lumens.",
              "minimum": 0,
              "type": "integer"
            },
            "sentAt": {
              "description": "Date and time when the message was sent.",
              "format": "date-time",
              "type": "string"
            }
          },
          "type": "object"
        },
        "summary": "Inform about environmental lighting conditions of a particular streetlight.",
        "title": "Light measured",
        "traits": [
          {
            "headers": {
              "properties": {
                "my-app-header": {
                  "maximum": 100,
                  "minimum": 0,
                  "type": "integer"
                }
              },
              "type": "object"
            }
          }
        ]
      },
      "turnOnOff": {
        "name": "turnOnOff",
        "payload": {
          "properties": {
            "command": {
              "description": "Whether to turn on or off the light.",
              "enum": [
                "on",
                "off"
              ],
              "type": "string"
            },
            "sentAt": {
              "description": "Date and time when the message was sent.",
              "format": "date-time",
              "type": "string"
            }
          },
          "type": "object"
        },
        "summary": "Command a particular streetlight to turn the lights on or off.",
        "title": "Turn on/off",
        "traits": [
          {
            "headers": {
              "properties": {
                "my-app-header": {
                  "maximum": 100,
                  "minimum": 0,
                  "type": "integer"
                }
              },
              "type": "object"
            }
          }
        ]
      }
    },
    "operationTraits": {
      "mqtt": {
        "bindings": {
          "mqtt": {
            "qos": 1
          }
        }
      }
    },
    "parameters": {
      "streetlightId": {
        "description": "The ID of the streetlight.",
        "schema": {
          "type": "string"
        }
      }
    },
    "schemas": {
      "dimLightPayload": {
        "properties": {
          "percentage": {
            "description": "Percentage to which the light should be dimmed to.",
            "maximum": 100,
            "minimum": 0,
            "type": "integer"
          },
          "sentAt": {
            "description": "Date and time when the message was sent.",
            "format": "date-time",
            "type": "string"
          }
        },
        "type": "object"
      },
      "lightMeasuredPayload": {
        "properties": {
          "lumens": {
            "description": "Light intensity measured in lumens.",
            "minimum": 0,
            "type": "integer"
          },
          "sentAt": {
            "description": "Date and time when the message was sent.",
            "format": "date-time",
            "type": "string"
          }
        },
        "type": "object"
      },
      "sentAt": {
        "description": "Date and time when the message was sent.",
        "format": "date-time",
        "type": "string"
      },
      "turnOnOffPayload": {
        "properties": {
          "command": {
            "description": "Whether to turn on or off the light.",
            "enum": [
              "on",
              "off"
            ],
            "type": "string"
          },
          "sentAt": {
            "description": "Date and time when the message was sent.",
            "format": "date-time",
            "type": "string"
          }
        },
        "type": "object"
      }
    },
    "securitySchemes": {
      "apiKey": {
        "description": "Provide your API key as the user and leave the password empty.",
        "in": "user",
        "type": "apiKey"
      },
      "openIdConnectWellKnown": {
        "openIdConnectUrl": "https://authserver.example/.well-known",
        "type": "openIdConnect"
      },
      "supportedOauthFlows": {
        "description": "Flows to support OAuth 2.0",
        "flows": {
          "authorizationCode": {
            "authorizationUrl": "https://authserver.example/auth",
            "refreshUrl": "https://authserver.example/refresh",
            "scopes": {
              "streetlights:dim": "Ability to dim the lights",
              "streetlights:off": "Ability to switch lights off",
              "streetlights:on": "Ability to switch lights on"
            },
            "tokenUrl": "https://authserver.example/token"
          },
          "clientCredentials": {
            "scopes": {
              "streetlights:dim": "Ability to dim the lights",
              "streetlights:off": "Ability to switch lights off",
              "streetlights:on": "Ability to switch lights on"
            },
            "tokenUrl": "https://authserver.example/token"
          },
          "implicit": {
            "authorizationUrl": "https://authserver.example/auth",
            "scopes": {
              "streetlights:dim": "Ability to dim the lights",
              "streetlights:off": "Ability to switch lights off",
              "streetlights:on": "Ability to switch lights on"
            }
          },
          "password": {
            "scopes": {
              "streetlights:dim": "Ability to dim the lights",
              "streetlights:off": "Ability to switch lights off",
              "streetlights:on": "Ability to switch lights on"
            },
            "tokenUrl": "https://authserver.example/token"
          }
        },
        "type": "oauth2"
      }
    }
  },
  "defaultContentType": "application/json",
  "info": {
    "description": "The Smartylighting Streetlights API allows you to remotely manage the city lights.\n### Check out its awesome features:\n* Turn a specific streetlight on/off 🌃\n* Dim a specific streetlight 😎\n* Receive real-time information about environmental lighting conditions 📈\n",
    "license": {
      "name": "Apache 2.0",
      "url": "https://www.apache.org/licenses/LICENSE-2.0"
    },
    "title": "Streetlights MQTT API",
    "version": "1.0.0"
  },
  "servers": {
    "production": {
      "description": "Test broker",
      "protocol": "mqtt",
      "security": [
        {
          "apiKey": []
        },
        {
          "supportedOauthFlows": [
            "streetlights:on",
            "streetlights:off",
            "streetlights:dim"
          ]
        },
        {
          "openIdConnectWellKnown": []
        }
      ],
      "url": "test.mosquitto.org:{port}",
      "variables": {	
        "port": {
          "default": "1883",
          "description": "Secure connection (TLS) is available through port 8883.",
          "enum": [
            "1883",
            "8883"
          ]
        }
      }
    }
  }
}
`)

func TestFoo(t *testing.T) {
	d := DocumentV2{}
	if err := json.Unmarshal(doc, &d); err != nil {
		t.Error(err)
		t.Fail()
	}
	fmt.Println(d)

}

func TestMapStructure(t *testing.T) {
	raw := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(doc, &raw))

	d := DocumentV2{}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		//DecodeHook:       nil,
		//ErrorUnused:      false,
		//ZeroFields:       false,
		//WeaklyTypedInput: false,
		//Squash:           false,
		//Metadata:         nil,
		Result:  &d,
		TagName: "json",
	})
	require.NoError(t, err)

	require.NoError(t, dec.Decode(raw))

	fmt.Println(d)
}

type Foo struct {
	Variables map[string]ServerVariableV2 `json:"variables"`
}

func TestMapStructure2(t *testing.T) {
	raw := make(map[string]interface{})
	require.NoError(t, json.Unmarshal(doc, &raw))

	var d DocumentV2

	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Squash: true,
		Result: &d,
	})
	require.NoError(t, err)
	require.NoError(t, dec.Decode(raw))

	fmt.Println(d.Servers())
	fmt.Println(d.Extension("x-test"))
}
