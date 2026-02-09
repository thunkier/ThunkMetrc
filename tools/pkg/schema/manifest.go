package schema

type ManifestEntry struct {
	ResponseType      string `json:"responseType"`
	SchemaFile        string `json:"schemaFile"`
	RequestType       string `json:"requestType,omitempty"`
	RequestSchemaFile string `json:"requestSchemaFile,omitempty"`
	IsRequestArray    bool   `json:"isRequestArray"`
	IsArray           bool   `json:"isArray"`
	IsPaginated       bool   `json:"isPaginated"`
	IsWrite           bool   `json:"isWrite"`
}
