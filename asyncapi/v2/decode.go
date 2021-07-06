package v2

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

// Decode implements the Decoder interface. Decodes AsyncAPI V2.x.x JSON Schema documents.
func Decode(b []byte, dst interface{}) error {
	raw := make(map[string]interface{})
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Squash: true,
		Result: dst,
	})

	if err != nil {
		return err
	}

	return dec.Decode(raw)
}
