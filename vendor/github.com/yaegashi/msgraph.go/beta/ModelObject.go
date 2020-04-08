// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ObjectDefinition undocumented
type ObjectDefinition struct {
	// Object is the base model of ObjectDefinition
	Object
	// Attributes undocumented
	Attributes []AttributeDefinition `json:"attributes,omitempty"`
	// Metadata undocumented
	Metadata []MetadataEntry `json:"metadata,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// SupportedApis undocumented
	SupportedApis []string `json:"supportedApis,omitempty"`
}

// ObjectIdentity undocumented
type ObjectIdentity struct {
	// Object is the base model of ObjectIdentity
	Object
	// SignInType undocumented
	SignInType *string `json:"signInType,omitempty"`
	// Issuer undocumented
	Issuer *string `json:"issuer,omitempty"`
	// IssuerAssignedID undocumented
	IssuerAssignedID *string `json:"issuerAssignedId,omitempty"`
}

// ObjectMapping undocumented
type ObjectMapping struct {
	// Object is the base model of ObjectMapping
	Object
	// AttributeMappings undocumented
	AttributeMappings []AttributeMapping `json:"attributeMappings,omitempty"`
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// FlowTypes undocumented
	FlowTypes *ObjectFlowTypes `json:"flowTypes,omitempty"`
	// Metadata undocumented
	Metadata []MetadataEntry `json:"metadata,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Scope undocumented
	Scope *Filter `json:"scope,omitempty"`
	// SourceObjectName undocumented
	SourceObjectName *string `json:"sourceObjectName,omitempty"`
	// TargetObjectName undocumented
	TargetObjectName *string `json:"targetObjectName,omitempty"`
}