// Code generated by github.com/johanneswuerbach/go-jsonschema, DO NOT EDIT.

package test

type Bar struct {
	// RefToFoo corresponds to the JSON schema field "refToFoo".
	RefToFoo *Foo `json:"refToFoo,omitempty" yaml:"refToFoo,omitempty" mapstructure:"refToFoo,omitempty"`
}

type Cyclic struct {
	// A corresponds to the JSON schema field "a".
	A *Foo `json:"a,omitempty" yaml:"a,omitempty" mapstructure:"a,omitempty"`
}

type Foo struct {
	// RefToBar corresponds to the JSON schema field "refToBar".
	RefToBar *Bar `json:"refToBar,omitempty" yaml:"refToBar,omitempty" mapstructure:"refToBar,omitempty"`
}
