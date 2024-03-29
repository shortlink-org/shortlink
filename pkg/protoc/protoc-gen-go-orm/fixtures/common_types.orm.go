// Code generated by protoc-gen-go-orm. DO NOT EDIT.
package fixtures

type StringFilterInput struct { // nolint:unused
	Eq          string   `json:"eq,omitempty"`
	Ne          string   `json:"ne,omitempty"`
	Lt          string   `json:"lt,omitempty"`
	Le          string   `json:"le,omitempty"`
	Gt          string   `json:"gt,omitempty"`
	Ge          string   `json:"ge,omitempty"`
	Contains    []string `json:"contains,omitempty"`
	NotContains []string `json:"notContains,omitempty"`
	StartsWith  string   `json:"startsWith,omitempty"`
	EndsWith    string   `json:"endsWith,omitempty"`
	IsEmpty     bool     `json:"isEmpty,omitempty"`
	IsNotEmpty  bool     `json:"isNotEmpty,omitempty"`
}
