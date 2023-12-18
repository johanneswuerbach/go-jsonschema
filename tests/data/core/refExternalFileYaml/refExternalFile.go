// Code generated by github.com/johanneswuerbach/go-jsonschema, DO NOT EDIT.

package test

type RefExternalFile struct {
	// MyExternalThing corresponds to the JSON schema field "myExternalThing".
	MyExternalThing *YamlStructNameFromFile `json:"myExternalThing,omitempty" yaml:"myExternalThing,omitempty" mapstructure:"myExternalThing,omitempty"`

	// SomeOtherExternalThing corresponds to the JSON schema field
	// "someOtherExternalThing".
	SomeOtherExternalThing *YamlStructNameFromFile `json:"someOtherExternalThing,omitempty" yaml:"someOtherExternalThing,omitempty" mapstructure:"someOtherExternalThing,omitempty"`
}

type YamlStructNameFromFile struct {
	// Foo corresponds to the JSON schema field "foo".
	Foo *string `json:"foo,omitempty" yaml:"foo,omitempty" mapstructure:"foo,omitempty"`
}
