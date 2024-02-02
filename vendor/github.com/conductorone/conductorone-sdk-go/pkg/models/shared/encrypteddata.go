// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EncryptedData is a message that contains encrypted bytes and metadata.
type EncryptedData struct {
	// The human-readable description of the encrypted data.
	Description *string `json:"description,omitempty"`
	// The encrypted bytes.
	EncryptedBytes *string `json:"encryptedBytes,omitempty"`
	// The key ID used to encrypt the data.
	KeyID *string `json:"keyId,omitempty"`
	// The human-readable name of the encrypted data.
	Name *string `json:"name,omitempty"`
	// The encryption provider used to encrypt the data.
	Provider *string `json:"provider,omitempty"`
	// The (optional) JSON schema of the encrypted data.
	Schema *string `json:"schema,omitempty"`
}

func (o *EncryptedData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *EncryptedData) GetEncryptedBytes() *string {
	if o == nil {
		return nil
	}
	return o.EncryptedBytes
}

func (o *EncryptedData) GetKeyID() *string {
	if o == nil {
		return nil
	}
	return o.KeyID
}

func (o *EncryptedData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EncryptedData) GetProvider() *string {
	if o == nil {
		return nil
	}
	return o.Provider
}

func (o *EncryptedData) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
