// Code generated by github.com/atombender/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import yaml "gopkg.in/yaml.v3"

type ObjectMyObject struct {
	// MyString corresponds to the JSON schema field "myString".
	MyString string `json:"myString" yaml:"myString" mapstructure:"myString"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ObjectMyObject) UnmarshalJSON(value []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(value, &raw); err != nil {
		return err
	}
	if v, ok := raw["myString"]; !ok || v == nil {
		return fmt.Errorf("field myString in ObjectMyObject: required")
	}
	type Plain ObjectMyObject
	var plain Plain
	if err := json.Unmarshal(value, &plain); err != nil {
		return err
	}
	*j = ObjectMyObject(plain)
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (j *ObjectMyObject) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}
	if v, ok := raw["myString"]; !ok || v == nil {
		return fmt.Errorf("field myString in ObjectMyObject: required")
	}
	type Plain ObjectMyObject
	var plain Plain
	if err := value.Decode(&plain); err != nil {
		return err
	}
	*j = ObjectMyObject(plain)
	return nil
}

type Object struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *ObjectMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}
