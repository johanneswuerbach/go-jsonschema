// Code generated by github.com/johanneswuerbach/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"
import "fmt"
import "time"

type DateTimeMyObject struct {
	// MyDateTime corresponds to the JSON schema field "myDateTime".
	MyDateTime time.Time `json:"myDateTime" yaml:"myDateTime" mapstructure:"myDateTime"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DateTimeMyObject) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["myDateTime"]; !ok || v == nil {
		return fmt.Errorf("field myDateTime in DateTimeMyObject: required")
	}
	type Plain DateTimeMyObject
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DateTimeMyObject(plain)
	return nil
}

type DateTime struct {
	// MyObject corresponds to the JSON schema field "myObject".
	MyObject *DateTimeMyObject `json:"myObject,omitempty" yaml:"myObject,omitempty" mapstructure:"myObject,omitempty"`
}
