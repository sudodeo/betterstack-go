package uptime

// Metadata represents the metadata information in the response.
type Metadata struct {
	Data       []MetadataRecord `json:"data,omitempty"`       // Data contains the metadata records.
	Pagination pagination       `json:"pagination,omitempty"` // Pagination contains the pagination information.
}

// MetadataRecord represents a metadata record in the response.
type MetadataRecord struct {
	ID            string             `json:"id,omitempty"`
	Type          string             `json:"type,omitempty"`
	Attributes    metadataAttributes `json:"attributes,omitempty"`
	Relationships relationships      `json:"relationships,omitempty"`
}

type metadataAttributes struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}
