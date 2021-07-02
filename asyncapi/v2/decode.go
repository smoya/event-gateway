package v2

import (
	"bytes"
	"encoding/json"
	"reflect"

	"github.com/asyncapi/event-gateway/asyncapi"
	"github.com/asyncapi/parser-go/pkg/parser"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// Decode implements the Decoder interface. Decodes AsyncAPI V2.x.x documents.
func Decode(b []byte, dst interface{}) error {
	r, err := parser.NewReader(string(b)) // parser should provide another method for parsing []byte
	if err != nil {
		return errors.Wrap(err, "error reading AsyncAPI doc")
	}

	p, err := parser.New()
	if err != nil {
		return err
	}

	w := bytes.NewBuffer(nil)
	if err := p(r, w); err != nil {
		return errors.Wrap(err, "error parsing AsyncAPI doc")
	}

	raw := make(map[string]interface{})
	if err := json.Unmarshal(w.Bytes(), &raw); err != nil {
		return err
	}

	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: setModelIdentifierHook,
		Squash:     true,
		Result:     dst,
	})
	if err != nil {
		return err
	}

	return dec.Decode(raw)
}

// setModelIdentifierHook is a hook for the mapstructure decoder.
// It checks if the destination field is a map of Identifiable elements and sets the proper identifier (name, id, etc) to it.
// Example: Useful for storing the name of the server in the Server struct (AsyncAPI doc does not have such field because it assumes the name is the key of the map).
func setModelIdentifierHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() != reflect.Map || to.Kind() != reflect.Map {
		return data, nil
	}

	identifiableInterface := reflect.TypeOf((*asyncapi.Identifiable)(nil)).Elem()
	if to.Key() != reflect.TypeOf("string") || !to.Elem().Implements(identifiableInterface) {
		return data, nil
	}

	fieldName := reflect.New(to.Elem()).Interface().(asyncapi.Identifiable).IDField()
	for k, v := range data.(map[string]interface{}) {
		// setting the value directly in the raw map. The struct needs to keep the mapstructure field tag so it unmarshals the field.
		v.(map[string]interface{})[fieldName] = k
	}

	return data, nil
}
