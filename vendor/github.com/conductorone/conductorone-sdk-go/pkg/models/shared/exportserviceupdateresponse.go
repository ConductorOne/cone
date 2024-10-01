// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ExportServiceUpdateResponse message.
type ExportServiceUpdateResponse struct {
	// The Exporter message.
	//
	// This message contains a oneof named export_to. Only a single field of the following list may be set at a time:
	//   - datasource
	//
	Exporter *Exporter `json:"exporter,omitempty"`
}

func (o *ExportServiceUpdateResponse) GetExporter() *Exporter {
	if o == nil {
		return nil
	}
	return o.Exporter
}
