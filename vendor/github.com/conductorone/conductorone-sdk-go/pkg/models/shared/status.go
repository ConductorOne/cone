// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// Details - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Details struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (d Details) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *Details) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Details) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *Details) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// Status - The `Status` type defines a logical error model that is suitable for
//
//	different programming environments, including REST APIs and RPC APIs. It is
//	used by [gRPC](https://github.com/grpc). Each `Status` message contains
//	three pieces of data: error code, error message, and error details.
//
//	You can find out more about this error model and how to work with it in the
//	[API Design Guide](https://cloud.google.com/apis/design/errors).
type Status struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code *int `json:"code,omitempty"`
	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	Details []Details `json:"details,omitempty"`
	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message *string `json:"message,omitempty"`
}

func (o *Status) GetCode() *int {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Status) GetDetails() []Details {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *Status) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}
